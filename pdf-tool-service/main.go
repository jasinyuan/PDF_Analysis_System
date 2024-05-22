package main

import (
	"log"
	routers "main/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	routers.Message(r)

	if err := r.Run(":8000"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}