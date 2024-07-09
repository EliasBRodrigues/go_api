package models

type Product struct {
	ID    int     `json:"id_product"` // estrutura para consumo rest
	Name  string  `json:"name"`       // nome de coluna
	Price float64 `json:"price"`
}
