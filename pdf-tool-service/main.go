package main

import (
	"log"
	routers "main/router"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	routers.Message(r)

	port := os.Getenv("PDF_SERVICE_PORT")
	if port == "" {
		port = "8000"
	}

	if err := r.Run(":8000"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}