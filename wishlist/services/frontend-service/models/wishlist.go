package models

import (
	"time"
)

type Wishlist struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	IsPublic     bool      `json:"is_public"`
	CanEdit      bool      `json:"can_edit"`
	CreatedAt    time.Time `json:"created_at"`
	LastModified time.Time `json:"last_modified"`
	LastOpened   time.Time `json:"last_opened"`
	Items        []Item    `json:"items"`
}
