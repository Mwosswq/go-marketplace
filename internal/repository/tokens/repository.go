package tokens

import (
	"context"
	"database/sql"
	"errors"
)

type Repository interface {
	CreateToken(ctx context.Context, id int, token string) error
	GetToken(ctx context.Context, id int) (string, error)
}

type repository struct {
	db *sql.DB
}

func (r *repository) CreateToken(ctx context.Context, id int, token string) error {
	queryRow := `INSERT INTO tokens (user_id, token) VALUES ($1, $2)`
	if _, err := r.db.ExecContext(ctx, queryRow, id, token); err != nil {
		return err
	}

	return nil
}

func (r *repository) GetToken(ctx context.Context, id int) (string, error) {
	var token string

	queryRow := `SELECT token from tokens where user_id=&1`
	if err := r.db.QueryRowContext(ctx, queryRow, id).Scan(&token); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	return token, nil
}
