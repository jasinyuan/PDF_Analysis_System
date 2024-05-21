package routes

import (
	uploadcor "businessservice/controllers/uploadCor"

	"github.com/gin-gonic/gin"
)

func Upload(r *gin.Engine) {
	r.GET("/upload", uploadcor.UploadControllers{}.UploadCor)
}