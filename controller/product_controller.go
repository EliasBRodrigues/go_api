package controller

import (
	"fmt"
	"go-api/models"
	"go-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productService service.ProductService
}

func NewProductController(service service.ProductService) productController {
	return productController{
		productService: service,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	// products := []models.Product{
	// 	{
	// 		ID:    2,
	// 		Name:  "batata",
	// 		Price: 10,
	// 	},
	// }
	products, error := p.productService.GetProducts()
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, error)
	}
	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var prod models.Product
	error := ctx.BindJSON(&prod)

	if error != nil {
		ctx.JSON(http.StatusBadRequest, error)
		return
	}

	insertedProductData, error := p.productService.CreateProduct(prod)
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, error)
	}

	ctx.JSON(http.StatusCreated, insertedProductData)
}

func (p *productController) GetProductsById(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		dto := models.DTO{
			Messsage: "ID can not be NULL",
		}
		ctx.JSON(http.StatusBadRequest, dto)
		return
	}

	prodId, error := strconv.Atoi(id)
	if error != nil {
		dto := models.DTO{
			Messsage: "ID need be a NUMBER",
		}
		ctx.JSON(http.StatusBadRequest, dto)
		return
	}

	product, error := p.productService.GetProductsById(prodId)
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, error)
		return
	}

	if product == nil {
		dto := models.DTO{
			Messsage: "Product NOT FOUND",
		}
		ctx.JSON(http.StatusNotFound, dto)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) UpdateProductById(ctx *gin.Context) {
	var prod models.Product
	error := ctx.BindJSON(&prod)

	if error != nil {
		ctx.JSON(http.StatusBadRequest, error)
		return
	}
	id := ctx.Param("prodID")
	prodID, error := strconv.Atoi(id)
	if error != nil {
		dto := models.DTO{Messsage: "ID need be a NUMBER"}
		ctx.JSON(http.StatusBadRequest, dto)
		return
	}

	updatedProductById, error := p.productService.UpdateProductById(prodID, prod)
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, error)
	}

	ctx.JSON(http.StatusOK, updatedProductById)
}

func (p *productController) DeleteProductById(ctx *gin.Context) {
	id := ctx.Param("prodID")
	prodID, error := strconv.Atoi(id)

	if error != nil {
		dto := models.DTO{Messsage: "ID need be a NUMBER"}
		ctx.JSON(http.StatusBadRequest, dto)
		return
	}

	product, error := p.productService.DeleteProductById(prodID)
	if error != nil {
		dto := models.DTO{Messsage: "FAILED"}
		ctx.JSON(http.StatusInternalServerError, dto)
		return
	}

	if product > 0 {
		dto := models.DTO{Messsage: fmt.Sprintf(`DELETED SUCCESS: %d`, prodID)}
		ctx.JSON(http.StatusOK, dto)
		return
	} else {
		dto := models.DTO{Messsage: fmt.Sprintf(`NOT FOUND: %d`, prodID)}
		ctx.JSON(http.StatusNotFound, dto)
		return
	}
}
