package main

import (
	"businessservice/routes"
	"net/http"
	"os"

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
	r.GET("/finish",func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "result.html", gin.H{
			"Message": "文件处理完成",
		})
	})
	r.GET("/file",func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "result.html", gin.H{
			"Message": "文件处理失败",
		})
	})

	// Start the server
	port := os.Getenv("BUSINESS_SERVICE_PORT")
	if port == "" {
		port = "8080"
	}
	r.Run()
}