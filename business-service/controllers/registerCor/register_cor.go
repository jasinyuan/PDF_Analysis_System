package registercor

import (
	"businessservice/model"
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
	user := &model.User{}
	if err := ctx.ShouldBind(&user);err == nil {
		ctx.JSON(http.StatusOK,user)
	}else{
		ctx.JSON(http.StatusOK,gin.H{
			"err":err.Error(),
		})
	}
}