package main

import (
	"go-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	ProductController := controller.NewProductController()

	server.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "running"})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":5050")
}
