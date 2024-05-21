package registercor

import (
	"businessservice/model"
	"businessservice/tool"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterControllers struct {
}

func (RegisterControllers) RegisterGetCor(ctx *gin.Context) {
	ctx.HTML(http.StatusOK,"register.html",gin.H{
		"title":"注册",
	})
}

func (RegisterControllers) RegisterPostCor(ctx *gin.Context) {
	personal := &model.User{}
	if err := ctx.ShouldBind(&personal);err == nil {
		tool.SearchDb(*personal,ctx,"register")
	}else{
		ctx.JSON(http.StatusOK,gin.H{
			"err":err.Error(),
		})
	}
}