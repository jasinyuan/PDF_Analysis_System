package routes

import (
	logincor "businessservice/controllers/loginCor"

	"github.com/gin-gonic/gin"
)

/* 登录请求路由 */
func Login(r *gin.Engine) {
	// 登录页面的HTML模板
	r.GET("/login",logincor.LogionControllers{}.LoginGetCor)
	// 通过表单获取登录的用户数据
	r.POST("/login",logincor.LogionControllers{}.LoginPostCor)
}