package auth

import (
	"context"
	"database/sql"
	"errors"
	"main/internal/domain"
)

type Repository interface {
	CreateUser(ctx context.Context, user domain.User) error
	GetUser(ctx context.Context, username string) (domain.User, error)
	CreateToken(ctx context.Context, id int, token string) error
}

type repository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user domain.User) error {
	queryStr := `insert into users (email, username, password) values ($1, $2, $3)`
	if _, err := r.db.ExecContext(ctx, queryStr, user.Email, user.Username, user.Password); err != nil {
		return err
	}

	return nil
}

func (r *repository) GetUser(ctx context.Context, username string) (domain.User, error) {
	var user domain.User

	queryStr := `select id, email, username, password from users where username=$1`
	if err := r.db.QueryRowContext(ctx, queryStr, username).Scan(&user.ID, &user.Email, &user.Username, &user.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, errors.New("user is not exists")
		}

		return domain.User{}, err
	}

	return user, nil
}

func (r *repository) CreateToken(ctx context.Context, id int, token string) error {
	queryStr := `insert into tokens (user_id, token) values ($1,$2)`
	if _, err := r.db.ExecContext(ctx, queryStr, id, token); err != nil {
		return err
	}

	return nil
}
