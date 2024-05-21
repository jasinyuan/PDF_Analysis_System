package routes

import (
	uploadcor "businessservice/controllers/uploadCor"

	"github.com/gin-gonic/gin"
)

func Upload(r *gin.Engine) {
	// 上传页面的HTML模板
	r.GET("/upload", uploadcor.UploadControllers{}.UploadGetCor)
	// 通过表单上传文件以及需求功能
	r.POST("/upload",uploadcor.UploadControllers{}.UploadPostCor)
}