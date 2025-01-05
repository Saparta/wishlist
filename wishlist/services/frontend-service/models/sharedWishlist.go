package models

import (
	"time"
)

type SharedWishlist struct {
	ID         string    `json:"id"`
	WishlistID string    `json:"wishlist_id"`
	SharedWith string    `json:"shared_with"`
	CanEdit    bool      `json:"can_edit"`
	CreatedAt  time.Time `json:"created_at"`
}
