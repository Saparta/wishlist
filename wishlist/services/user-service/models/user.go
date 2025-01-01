package models

import "time"

type User struct {
	ID        string    `json:"id"`
	OAuthID   string    `json:"oauth_id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
