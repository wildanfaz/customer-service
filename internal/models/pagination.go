package models

type (
	Pagination struct {
		Page    int `json:"-" query:"page"`
		PerPage int `json:"-" query:"per_page"`
	}
)