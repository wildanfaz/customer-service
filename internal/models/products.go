package models

import "time"

type (
	Product struct {
		Id          int       `json:"id"`
		Name        string    `json:"name" query:"name"`
		Description string    `json:"description" query:"description"`
		Quantity    int       `json:"quantity"`
		Price       int       `json:"price"`
		Category    string    `json:"category" query:"category"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		Pagination
	}

	Products []Product
)