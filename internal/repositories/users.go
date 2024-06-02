package repositories

import (
	"context"
	"database/sql"

	"github.com/wildanfaz/go-market/internal/models"
)

type ImplementUsers struct {
	db *sql.DB
}

type Users interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	Register(ctx context.Context, payload models.User) error
}

func NewUsersRepository(db *sql.DB) Users {
	return &ImplementUsers{
		db: db,
	}
}

func (r *ImplementUsers) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user := models.User{}

	query := `
	SELECT id, full_name, email, password, balance, created_at, updated_at
	FROM users
	WHERE email = ?
	`

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.Id, &user.FullName, &user.Email, &user.Password, &user.Balance, &user.CreatedAt, &user.UpdatedAt,
	)

	return &user, err
}

func (r *ImplementUsers) Register(ctx context.Context, payload models.User) error {
	query := `INSERT INTO users (full_name, email, password) VALUES (?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, payload.FullName, payload.Email, payload.Password)

	return err
}