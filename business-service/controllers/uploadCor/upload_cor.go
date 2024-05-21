package uploadcor

import (
	"io/ioutil"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type UploadControllers struct {
}

/* 上传页面的HTML模板 */
func (UploadControllers) UploadGetCor(ctx *gin.Context) {
	ctx.HTML(http.StatusOK,"upload.html",gin.H{
		"title":"上传",
	})
}

/* 上传文件发布消息 */
func (UploadControllers) UploadPostCor(ctx *gin.Context, rdb *redis.Client) {
	// 获取传入的文件
	file,err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"err":"文件获取失败",
		})
		return
	}

	// 获取需求功能
	demand,_ := ctx.FormFile("demand")

	// 设置上传至本地的路径
	dst := path.Join("./static",file.Filename)

	if err := ctx.SaveUploadedFile(file,dst); err != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"err":"文件上传至本地失败",
		})
		return
	}

	// 读取二进制文件内容
    fileContent, err := ioutil.ReadFile(dst)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": "文件解析失败",
		})
        return
    }

	// 组合消息文件内容
	message := map[string]interface{}{
		"demand":demand,
		"fileName":file.Filename,
		"fileData":fileContent,
	}

	// redis发布信息
	err = rdb.Publish(ctx, "pdf_channel", message).Err()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": "pdf消息发布失败",
		})
        return
	}

}
