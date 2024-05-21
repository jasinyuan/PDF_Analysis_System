package logincor

import (
	"businessservice/model"
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
	user := &model.User{}
	if err := ctx.ShouldBind(&user); err == nil {
		ctx.JSON(http.StatusOK,user)
	}else{
		ctx.JSON(http.StatusOK,gin.H{
			"err":err.Error(),
		})
	}
}