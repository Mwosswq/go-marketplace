package items

import (
	"context"
	"database/sql"
	"items-service/internal/domain"
)

type Repository interface {
	CreateItem(ctx context.Context, item *domain.Item) (int32, error)
	GetItem(ctx context.Context, id int32) (domain.Item, error)
}

type repository struct {
	db *sql.DB
}

func NewItemsRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateItem(ctx context.Context, item *domain.Item) (int32, error) {
	var id int32

	queryStr := "INSERT INTO items (title, description, created_at, price) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.db.QueryRowContext(ctx, queryStr, item.Title, item.Description, item.CreatedAt, item.Price).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, err
}

func (r *repository) GetItem(ctx context.Context, id int32) (domain.Item, error) {
	var item domain.Item

	queryStr := "SELECT id, title, description, created_at, price FROM items WHERE id=$1"
	err := r.db.QueryRowContext(ctx, queryStr, id).Scan(
		&item.ID,
		&item.Title,
		&item.Description,
		&item.CreatedAt,
		&item.Price,
	)

	if err != nil {
		return item, err
	}

	return item, nil
}
