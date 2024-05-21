package uploadcor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UploadControllers struct {
}

func (UploadControllers) UploadCor(ctx *gin.Context) {
	ctx.HTML(http.StatusOK,"upload.html",gin.H{
		"title":"上传",
	})
}