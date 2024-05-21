package routes

import (
	registercor "businessservice/controllers/registerCor"

	"github.com/gin-gonic/gin"
)

/* 注册请求路由 */
func Register(r *gin.Engine) {
	// 注册页面的HTML模板
	r.GET("/register",registercor.RegisterControllers{}.RegisterGetCor)
	// 通过表单获取注册的用户数据
	r.POST("/register",registercor.RegisterControllers{}.RegisterPostCor)
}