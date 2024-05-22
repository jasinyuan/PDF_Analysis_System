package massagecor

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"main/model"
	"main/sercives"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type MassageControllers struct {
	rdb *redis.Client
}

func NewMassageControllers(rdb *redis.Client) *MassageControllers {
    return &MassageControllers{rdb: rdb}
}

func (ctrl MassageControllers) MessageGet(ctx *gin.Context) {
	messageChan := make(chan string)
	go func() {
		message := subMessage(ctx)
		messageChan <- message
		close(messageChan)
	}()

	message := <- messageChan
	err := ctrl.rdb.Publish(ctx,"finish_channel",message).Err()
	if err != nil {
		log.Printf("消息发送至redis失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": "消息发布失败",
		})
        return
	}else{
		ctx.JSON(http.StatusOK,gin.H{
			"消息获取":message,
		})
	}
}

/* 订阅消息 */
func subMessage(ctx *gin.Context) string {
	rdb := model.GetRedisClient()
	subMessage := rdb.Subscribe(ctx,"pdf_channel")
	defer subMessage.Close()

	ch := subMessage.Channel()
	for msg := range ch {
		pdf := model.Pdf{}
		err := json.Unmarshal([]byte(msg.Payload), &pdf)
        if err != nil {
            log.Println("Error unmarshalling message:", err)
            continue
        }

		// Base64 解码文件数据
		fileData, err := base64.StdEncoding.DecodeString(pdf.FileData)
		if err != nil {
			log.Println("Error decoding base64 file data:", err)
			continue
		}
		dst := path.Join("./static", pdf.FileName)
		err1 := ioutil.WriteFile(dst, fileData, 0644)
		if err1 != nil {
			log.Println("Error writing file:", err)
			continue
		}

		switch pdf.Demand{
		case "split":
			err = sercives.SplitPDF(dst)
			if err != nil {
				log.Println("Error splitting PDF file:", err)
				return "fail"
			}
			log.Println("PDF 文件处理并成功拆分:", pdf.FileName)
			return "finish"
		case "merge":
			fmt.Println("合并功能")
		case "watermark":
			fmt.Println("加水印功能")
		default:
			fmt.Println("未知功能")
		}
	}
	return "file"
}