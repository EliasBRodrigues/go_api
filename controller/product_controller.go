package controller

import (
	"go-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
}

func NewProductController() productController {
	return productController{}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []models.Product{
		{
			ID:    2,
			Name:  "batata",
			Price: 10,
		},
	}
	ctx.JSON(http.StatusOK, products)
}
