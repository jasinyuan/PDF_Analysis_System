package uploadcor

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type UploadControllers struct {
	rdb *redis.Client
}

func NewUploadController(rdb *redis.Client) *UploadControllers {
    return &UploadControllers{rdb: rdb}
}

/* 上传页面的HTML模板 */
func (UploadControllers) UploadGetCor(ctx *gin.Context) {
	ctx.HTML(http.StatusOK,"upload.html",gin.H{
		"title":"上传",
	})
}

/* 上传文件发布消息 */
func (ctrl UploadControllers) UploadPostCor(ctx *gin.Context) {
	// 获取传入的文件
	file,err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"err":"文件获取失败",
		})
		return
	}

	// 获取需求功能
	demand := ctx.PostForm("demand")

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

	encodedFileContent := base64.StdEncoding.EncodeToString(fileContent)

	// 组合消息文件内容
	message := map[string]interface{}{
		"demand":demand,
		"fileName":file.Filename,
		"fileData":encodedFileContent,
	}

	messageJSON,err := json.Marshal(message)
	if err != nil {
		log.Printf("序列化消息失败: %v", err)
		return
	}

	// redis发布信息
	err = ctrl.rdb.Publish(ctx, "pdf_channel", messageJSON).Err()
	if err != nil {
		log.Printf("消息发送至redis失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": "pdf消息发布失败",
		})
        return
	}

	// 等待处理完成的消息
	responseChan := make(chan string)
	go func() {
		message := subscribeChannel(ctx, ctrl.rdb)
		responseChan <- message
		close(responseChan)
	}()

	select {
	case response := <-responseChan:
		if response == "finish" {
			ctx.Redirect(http.StatusMovedPermanently,"/finish")
		}else{
			ctx.Redirect(http.StatusMovedPermanently,"/file")
		}
	case <-time.After(15 * time.Second):
		ctx.JSON(http.StatusGatewayTimeout, gin.H{"err": "等待处理消息超时"})
	}
}

func subscribeChannel(ctx context.Context, rdb *redis.Client,) string{
	pubsub := rdb.Subscribe(ctx, "finish_channel")
	defer pubsub.Close()

	ch := pubsub.Channel()

	for msg := range ch {
		return msg.Payload
	}
	return "flie"
}
