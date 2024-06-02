package models

import "time"

type (
	Cart struct {
		UserId    int       `json:"-"`
		ProductId int       `json:"product_id"`
		Quantity  int       `json:"quantity"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)