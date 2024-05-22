package sercives

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

func SplitPDF(filePath string) error {
    // 创建奇偶页输出目录
    outputDir := path.Dir(filePath)
    oddDir := filepath.Join(outputDir, "odd_pages")
    evenDir := filepath.Join(outputDir, "even_pages")

    err := os.MkdirAll(oddDir, 0755)
    if err != nil {
        return fmt.Errorf("failed to create odd pages directory: %w", err)
    }
    err = os.MkdirAll(evenDir, 0755)
    if err != nil {
        return fmt.Errorf("failed to create even pages directory: %w", err)
    }

    // 使用 pdfcpu 拆分 PDF 文件，每页拆分成一个文件
    cmd := exec.Command("pdfcpu", "split", filePath, outputDir, "1")
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("failed to split PDF file: %s, error: %w", string(output), err)
    }

    // 移动奇偶页文件到对应目录
    files, err := os.ReadDir(outputDir)
    if err != nil {
        return fmt.Errorf("failed to read directory: %w", err)
    }

    type pdfPage struct {
        path string
        page int
    }

    oddFiles := []pdfPage{}
    evenFiles := []pdfPage{}

    for _, file := range files {
        if file.IsDir() || !strings.HasPrefix(file.Name(), "lryh_") {
            continue
        }
        srcPath := filepath.Join(outputDir, file.Name())
        var pageNum int
        _, err := fmt.Sscanf(file.Name(), "lryh_%d.pdf", &pageNum)
        if err != nil {
            return fmt.Errorf("failed to parse page number from file name: %w", err)
        }
        if pageNum%2 == 1 {
            oddFiles = append(oddFiles, pdfPage{srcPath, pageNum})
        } else {
            evenFiles = append(evenFiles, pdfPage{srcPath, pageNum})
        }
    }

    // 对文件按页码进行排序
    sort.Slice(oddFiles, func(i, j int) bool { return oddFiles[i].page < oddFiles[j].page })
    sort.Slice(evenFiles, func(i, j int) bool { return evenFiles[i].page < evenFiles[j].page })

    oddFilePaths := make([]string, len(oddFiles))
    for i, f := range oddFiles {
        oddFilePaths[i] = f.path
    }

    evenFilePaths := make([]string, len(evenFiles))
    for i, f := range evenFiles {
        evenFilePaths[i] = f.path
    }

    // 合并奇数页文件
    oddOutput := filepath.Join(oddDir, "odd_pages.pdf")
    err = mergePDFs(oddFilePaths, oddOutput)
    if err != nil {
        return fmt.Errorf("failed to merge odd pages: %w", err)
    }

    // 合并偶数页文件
    evenOutput := filepath.Join(evenDir, "even_pages.pdf")
    err = mergePDFs(evenFilePaths, evenOutput)
    if err != nil {
        return fmt.Errorf("failed to merge even pages: %w", err)
    }

    // 清理单页文件
    for _, file := range append(oddFilePaths, evenFilePaths...) {
        os.Remove(file)
    }

    return nil
}


func mergePDFs(files []string, output string) error {
    args := append([]string{"merge", output}, files...)
    cmd := exec.Command("pdfcpu", args...)
    outputBytes, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("failed to merge PDF files: %s, error: %w", string(outputBytes), err)
    }
    return nil
}
