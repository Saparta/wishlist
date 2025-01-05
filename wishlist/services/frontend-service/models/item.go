package models

import (
	"time"
)

type Item struct {
	ID         string    `json:"id"`
	WishlistID string    `json:"wishlist_id"`
	Name       string    `json:"name"`
	Url        string    `json:"url"`
	Price      float32   `json:"price"`
	IsGifted   bool      `json:"is_gifted"`
	GiftedBy   string    `json:"gifted_by"`
	CreatedAt  time.Time `json:"created_at"`
}