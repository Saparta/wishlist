package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        string `db:"id" json:"id"`
	Email     string `db:"email" json:"email"`
	Name      string `db:"name" json:"name"`
	AvatarURL string `db:"avatar_url" json:"avatar_url"`
}

type Store struct{ DB *sql.DB }

func (s *Store) FindOrCreateUserByEmail(ctx context.Context, email, name, avatar string) (*User, error) {
	u := &User{}
	err := s.DB.QueryRowContext(ctx, `
		INSERT INTO users (email, name, avatar_url)
		VALUES ($1,$2,$3)
		ON CONFLICT (email) DO UPDATE SET name=EXCLUDED.name, avatar_url=EXCLUDED.avatar_url
		RETURNING id, email, name, avatar_url
	`, email, name, avatar).Scan(&u.ID, &u.Email, &u.Name, &u.AvatarURL)
	return u, err
}

func (s *Store) GetUserByID(ctx context.Context, id string) (*User, error) {
	u := &User{}
	err := s.DB.QueryRowContext(ctx, `
		SELECT id, email, name, avatar_url FROM users WHERE id=$1
	`, id).Scan(&u.ID, &u.Email, &u.Name, &u.AvatarURL)
	return u, err
}