package migrations

import "database/sql"

func AddBalance(db *sql.DB, email string) error {
	query := `UPDATE users SET balance = balance + 1000000000 WHERE email = ?`

	_, err := db.Exec(query, email)

	return err
}