package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* 注册的间接路由页面 */
func Pop_success(r *gin.Engine)  {
	// 注册成功跳转页面
	r.GET("/pop-ups",func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK,"pop-ups.html",gin.H{
		})
	})
}