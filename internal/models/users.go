package models

import "time"

type (
	User struct {
		Id        int       `json:"id"`
		FullName  string    `json:"full_name"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		Balance   int       `json:"balance"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)