package items

import (
	"context"
	"database/sql"
	"main/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, item *domain.Item) error
	GetAllItems(ctx context.Context) ([]domain.Item, error)
	RemoveItem(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func NewItemsRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, item *domain.Item) error {
	queryStr := "INSERT INTO items (title, description, created_at, price) VALUES ($1, $2, $3, $4)"
	if _, err := r.db.ExecContext(ctx, queryStr, item.Title, item.Description, item.CreatedAt, item.Price); err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAllItems(ctx context.Context) ([]domain.Item, error) {
	var items []domain.Item

	queryStr := "SELECT ID, title, description, created_at, price FROM items"
	rows, err := r.db.QueryContext(ctx, queryStr)

	if err != nil {
		return items, err
	}

	defer rows.Close()

	for rows.Next() {
		var i domain.Item
		err := rows.Scan(&i.ID, &i.Title, &i.Description, &i.CreatedAt, &i.Price)
		if err != nil {
			return items, err
		}
		items = append(items, i)
	}
	return items, rows.Err()
}

func (r *repository) RemoveItem(ctx context.Context, id int) error {
	queryStr := "DELETE FROM items WHERE id=$1"
	_, err := r.db.ExecContext(ctx, queryStr, id)

	return err
}
