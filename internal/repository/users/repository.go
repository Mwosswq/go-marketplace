package users

import (
	"context"
	"database/sql"
	"errors"
	"main/internal/domain"
)

type Repository interface {
	GetUserByID(ctx context.Context, id int) (domain.User, error)
}

type repository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetUserByID(ctx context.Context, id int) (domain.User, error) {
	var user domain.User

	queryRow := `select id, username, email from users where id=$1`
	err := r.db.QueryRowContext(ctx, queryRow, id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}

	return user, nil
}
