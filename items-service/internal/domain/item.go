package domain

import (
	"time"
)

type Item struct {
	ID          int32     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	Price       float64   `json:"price"`
}
