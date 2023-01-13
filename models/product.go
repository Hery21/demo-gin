package models

type Product struct {
	ID          int    `json:"id"` gormescription string `json:"description"`
	Quantity    int    `json:"quantity"`
}
