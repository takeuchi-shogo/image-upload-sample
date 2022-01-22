package main

import (
	"golang-image-upload/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// config := config.Config.NewConfig()
	r := gin.Default()

	// r.GET("/images", handler.Download)
	r.POST("/images", handler.Uploader)

	r.Run(":8080")
}
