package models

import (
	"time"
)

type SharedWishlist struct {
	ID         string
	WishlistID string
	SharedWith string
	CanEdit    bool
	CreatedAt  time.Time
}
