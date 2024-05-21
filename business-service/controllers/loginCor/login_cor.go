package logincor

import (
	"businessservice/model"
	"businessservice/tool"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogionControllers struct {

}


func ( LogionControllers) LoginGetCor(ctx *gin.Context) {
	ctx.HTML(http.StatusOK,"login.html",gin.H{
		"title":"登录",
	})
}

func ( LogionControllers) LoginPostCor(ctx *gin.Context) {
	personal := &model.User{}
	if err := ctx.ShouldBind(&personal); err == nil {
		tool.SearchDb(*personal,ctx,"login")
	}else{
		ctx.JSON(http.StatusOK,gin.H{
			"err":err.Error(),
		})
	}
}
