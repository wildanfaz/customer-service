package repositories

import (
	"context"
	"database/sql"
	"strings"

	"github.com/wildanfaz/go-market/internal/helpers"
	"github.com/wildanfaz/go-market/internal/models"
)

type ImplementProducts struct {
	db *sql.DB
}

type Products interface{
	ListProducts(ctx context.Context, payload models.Product) (*models.Products, error)
}

func NewProductsRepository(db *sql.DB) Products {
	return &ImplementProducts{
		db: db,
	}
}

func (r *ImplementProducts) ListProducts(ctx context.Context, payload models.Product) (*models.Products, error) {
	values := []any{}
	search := ""
	query := `
	SELECT id, name, description, quantity, price, category, created_at, updated_at
	FROM products
	`

	if !helpers.IsZeroStruct(payload) {
		query += `WHERE MATCH(name, description, category) AGAINST(?)
		`
		
		if payload.Name != "" {
			search += payload.Name + " "
		}

		if payload.Description != "" {
			search += payload.Description + " "
		}

		if payload.Category != "" {
			search += payload.Category + " "
		}

		query = strings.TrimSuffix(query, " ")
		values = append(values, search)
	}

	query += `ORDER BY created_at DESC, name
	`

	limit, offset := 10, 0

	if payload.Pagination.PerPage > 0 {
		limit = payload.Pagination.PerPage
	}

	if payload.Pagination.Page > 1 {
		offset = (payload.Pagination.Page - 1) * limit
	}

	query += `LIMIT ? OFFSET ?
	`
	values = append(values, limit, offset)

	rows, err := r.db.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := models.Products{}

	for rows.Next() {
		product := models.Product{}

		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Quantity,
			&product.Price,
			&product.Category,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return &products, nil
}