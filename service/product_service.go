package service

import (
	"go-api/models"
	"go-api/repository"
)

type ProductService struct {
	repository repository.ProductRepository
}

func NewProductService(repos repository.ProductRepository) ProductService {
	return ProductService{
		repository: repos,
	}
}

func (ps *ProductService) GetProducts() ([]models.Product, error) {
	return ps.repository.GetProducts()
}

func (ps *ProductService) CreateProduct(prod models.Product) (models.Product, error) {
	prodId, error := ps.repository.CreateProduct(prod)
	if error != nil {
		return models.Product{}, error
	}

	prod.ID = prodId
	return prod, nil
}

func (ps *ProductService) GetProductsById(id_prod int) (*models.Product, error) {
	prod, error := ps.repository.GetProductsById(id_prod)
	if error != nil {
		return nil, error
	}

	return prod, nil
}

func (ps *ProductService) UpdateProductById(id int, body models.Product) (models.Product, error) {
	prod, error := ps.repository.UpdateProductById(id, body)
	if error != nil {
		return models.Product{}, error
	}
	body.ID = prod
	return body, error
}

func (ps *ProductService) DeleteProductById(id int) (int, error) {
	prod, error := ps.repository.DeleteProductById(id)
	if error != nil {
		return 0, error
	}
	return prod, error
}
