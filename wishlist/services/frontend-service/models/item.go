package models

import (
	"time"
)

type Item struct {
	WishlistID string    `json:"wishlist_id"`
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Url        string    `json:"url"`
	Price      float32   `json:"price"`
	IsGifted   bool      `json:"is_gifted"`
	GiftedBy   string    `json:"gifted_by"`
	CreatedAt  time.Time `json:"created_at"`
}
