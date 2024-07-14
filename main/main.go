package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/service"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnect, error := db.ConnectDB()
	if error != nil {
		panic(error)
	}

	ProductRepository := repository.NewProductRepository(dbConnect)
	ProductService := service.NewProductService(ProductRepository)
	ProductController := controller.NewProductController(ProductService)

	server.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "running"})
	})

	server.GET("/products", ProductController.GetProducts)
	server.GET("/product/:productId", ProductController.GetProductsById)
	server.POST("/products", ProductController.CreateProduct)
	server.PUT("/products/:prodID", ProductController.UpdateProductById)
	server.DELETE("/product/:prodID", ProductController.DeleteProductById)

	server.Run(":8080")
}
