package main

import (
	"businessservice/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 设置html模板
	r.Static("/static","./templates")
	r.LoadHTMLGlob("templates/*")
	// 路由设置
	routes.Upload(r)
	routes.Login(r)
	routes.Register(r)
	routes.Register_success(r)
	routes.Pop_success(r)
	// 将页面默认路由重定向
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/login")
	})
	r.Run()
}