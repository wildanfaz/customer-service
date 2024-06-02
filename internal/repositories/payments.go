package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/wildanfaz/go-market/internal/constants"
	"github.com/wildanfaz/go-market/internal/models"
)

type ImplementPayments struct {
	db *sql.DB
}

type Payments interface {
	Pay(ctx context.Context, payload models.User) error
}

func NewPaymentsRepository(db *sql.DB) Payments {
	return &ImplementPayments{
		db: db,
	}
}

func (r *ImplementPayments) Pay(ctx context.Context, payload models.User) error {
	query := `
	SELECT SUM(p.price * c.quantity) AS total_price
	FROM carts c
	JOIN products p ON p.id = c.product_id
	JOIN users u ON u.id = c.user_id
	WHERE u.id = ? AND c.is_purchased = 0
	FOR UPDATE
	`

	var totalPrice *int

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})
	if err != nil {
		return err
	}

	err = tx.QueryRowContext(ctx, query, payload.Id).Scan(&totalPrice)
	if err != nil {
		tx.Rollback()
		return err
	}

	if totalPrice == nil {
		tx.Rollback()
		return errors.New(constants.EmptyCart)
	}

	if payload.Balance < *totalPrice {
		tx.Rollback()
		return errors.New(constants.InsufficientBalance)
	}

	query = `UPDATE users SET balance = balance - ? WHERE id = ?`

	_, err = tx.ExecContext(ctx, query, totalPrice, payload.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = `UPDATE carts SET is_purchased = 1 WHERE user_id = ?`

	_, err = tx.ExecContext(ctx, query, payload.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}