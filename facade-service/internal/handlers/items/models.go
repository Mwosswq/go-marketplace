package items

import "time"

type CreatingItemsResponse struct {
	Id int32 `json:"id"`
}

type CreatingItemsRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type GettingItemsResponse struct {
	Id          int32     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	Price       float64   `json:"price"`
}

type GettingItemsRequest struct {
	Id int32 `json:"id"`
}
