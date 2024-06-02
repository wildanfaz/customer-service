package repositories

import (
	"context"
	"database/sql"

	"github.com/wildanfaz/go-market/internal/models"
)

type ImplementCarts struct {
	db *sql.DB
}

type Carts interface {
	AddToCart(ctx context.Context, payload models.Cart) error
	ListInCart(ctx context.Context, userId int) (*models.Products, error)
	DeleteFromCart(ctx context.Context, userId, productId int) error
}

func NewCartsRepository(db *sql.DB) Carts {
	return &ImplementCarts{
		db: db,
	}
}

func (r *ImplementCarts) AddToCart(ctx context.Context, payload models.Cart) error {
	query := `INSERT INTO carts (user_id, product_id, quantity) VALUES (?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, payload.UserId, payload.ProductId, payload.Quantity)

	return err
}

func (r *ImplementCarts) ListInCart(ctx context.Context, userId int) (*models.Products, error) {
	query := `
	SELECT 
	c.id, p.name, p.description, 
	c.quantity, (p.price * c.quantity) AS total_price, p.category, 
	c.created_at, c.updated_at
	FROM products p
	JOIN carts c ON p.id = c.product_id
	JOIN users u ON u.id = c.user_id
	WHERE u.id = ? AND c.is_purchased = 0
	`

	products := models.Products{}

	rows, err := r.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

func (r *ImplementCarts) DeleteFromCart(ctx context.Context, userId, productId int) error {
	query := `DELETE FROM carts WHERE user_id = ? AND product_id = ?`

	res, err := r.db.ExecContext(ctx, query, userId, productId)
	if err != nil {
		return err
	}

	af, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if af < 1 {
		return sql.ErrNoRows
	}

	return nil
}