package models

import (
	"time"
)

type Item struct {
	ID         string
	WishlistID string
	Name       string
	Url        string
	Price      float32
	IsGifted   bool
	GiftedBy   string
	CreatedAt  time.Time
}
