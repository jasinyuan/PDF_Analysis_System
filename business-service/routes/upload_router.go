package routes

import (
	uploadcor "businessservice/controllers/uploadCor"
	"businessservice/model"

	"github.com/gin-gonic/gin"
)

func Upload(r *gin.Engine) {
	// 获取redis
	rdb := model.GetRedisClient()
	// 实例化 UploadControllers
	UploadControllers := uploadcor.NewUploadController(rdb)
	// 上传页面的HTML模板
	r.GET("/upload", UploadControllers.UploadGetCor)
	// 通过表单上传文件以及需求功能
	r.POST("/upload",UploadControllers.UploadPostCor)
}