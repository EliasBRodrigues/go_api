package repository

import (
	"database/sql"
	"fmt"
	"go-api/models"
)

type ProductRepository struct {
	conn *sql.DB
}

func NewProductRepository(conn *sql.DB) ProductRepository {
	return ProductRepository{
		conn: conn,
	}
}

func (pr *ProductRepository) GetProducts() ([]models.Product, error) {
	query := "SELECT id, product_name, price FROM product ORDER BY id ASC"
	rows, error := pr.conn.Query(query)
	if error != nil {
		fmt.Println(error)
		return []models.Product{}, error
	}

	var productList []models.Product
	var productObject models.Product

	for rows.Next() {
		error = rows.Scan(
			&productObject.ID,
			&productObject.Name,
			&productObject.Price)
		if error != nil {
			fmt.Println(error)
			return []models.Product{}, error
		}
		productList = append(productList, productObject)
	}
	rows.Close()
	return productList, nil
}

func (pr *ProductRepository) CreateProduct(prod models.Product) (int, error) {
	var id int
	query, error := pr.conn.Prepare("INSERT INTO product" +
		"(product_name, price)" +
		"VALUES ($1, $2) RETURNING id")
	if error != nil {
		fmt.Println(error)
		return 0, error
	}

	error = query.QueryRow(prod.Name, prod.Price).Scan(&id)
	if error != nil {
		fmt.Println(error)
		return 0, error
	}
	query.Close()
	return id, nil
}

func (pr *ProductRepository) GetProductsById(id_prod int) (*models.Product, error) {
	query, error := pr.conn.Prepare("SELECT * FROM product WHERE id = $1")
	if error != nil {
		fmt.Println(error)
		return nil, error
	}

	var prod models.Product

	error = query.QueryRow(id_prod).Scan(
		&prod.ID,
		&prod.Name,
		&prod.Price,
	)

	if error != nil {
		if error == sql.ErrNoRows {
			return nil, nil
		}

		return nil, error
	}

	query.Close()
	return &prod, nil
}

func (pr *ProductRepository) UpdateProductById(id int, body models.Product) (int, error) {
	query, _ := pr.conn.Prepare(`UPDATE product SET product_name = $1, price = $2 WHERE id = $3 RETURNING id`)

	var updatedID int

	err := query.QueryRow(body.Name, body.Price, id).Scan(&updatedID)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	query.Close()
	return updatedID, nil

}

func (pr *ProductRepository) DeleteProductById(id_prod int) (int, error) {
	query := `DELETE FROM product WHERE id = $1`
	_, error := pr.conn.Exec(query, id_prod)
	if error != nil {
		fmt.Println(error)
		return 0, error
	}

	return id_prod, nil
}
