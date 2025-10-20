package users

import (
	"context"
	"database/sql"
	"errors"
	"main/internal/domain"
)

type Repository interface {
	CreateUser(ctx context.Context, user domain.User) error
	CheckEmailExisting(ctx context.Context, email string) (bool, error)
	CheckUsernameExisting(ctx context.Context, username string) (bool, error)
	GetUser(ctx context.Context, id int) (domain.UserResponse, error)
	GetUserForLogin(ctx context.Context, username string) (domain.User, error)
}

type repository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user domain.User) error {

	queryStr := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) "
	if _, err := r.db.ExecContext(ctx, queryStr, user.Username, user.Email, user.Password); err != nil {
		return err
	}

	return nil
}

func (r *repository) CheckEmailExisting(ctx context.Context, email string) (bool, error) {
	var exists bool

	queryRow := `SELECT exists(select 1 from users where email=$1)`
	if err := r.db.QueryRowContext(ctx, queryRow, email).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (r *repository) CheckUsernameExisting(ctx context.Context, username string) (bool, error) {
	var exists bool

	queryRow := `SELECT exists(select 1 from users where username=$1)`
	if err := r.db.QueryRowContext(ctx, queryRow, username).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (r *repository) GetUser(ctx context.Context, id int) (domain.UserResponse, error) {
	var user domain.UserResponse

	queryRow := `select id, username, email from users where id=$1`
	err := r.db.QueryRowContext(ctx, queryRow, id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.UserResponse{}, errors.New("user not found")
		}
		return domain.UserResponse{}, err
	}

	return user, nil
}

func (r *repository) GetUserForLogin(ctx context.Context, username string) (domain.User, error) {
	var user domain.User

	queryRow := `SELECT id, username, email, password FROM users WHERE username=$1`
	if err := r.db.QueryRowContext(ctx, queryRow, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}

	return user, nil
}
