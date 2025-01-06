package models

import (
	"time"
)

type Wishlist struct {
	ID           string
	UserID       string
	Title        string
	Description  string
	IsPublic     bool
	CreatedAt    time.Time
	LastModified time.Time
	LastOpened   time.Time
}
