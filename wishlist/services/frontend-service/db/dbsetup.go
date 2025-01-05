package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func createUsersTable(dbpool *pgxpool.Pool) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY,
		oauth_id UUID NOT NULL,
		email TEXT NOT NULL,
		name TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`

	if _, err := dbpool.Exec(context.Background(), query); err != nil {
		return err
	}

	return nil
}

func SetUpDb(channel chan *pgxpool.Pool) {
	var dsn string = os.Getenv("DSN")

	dbPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	if err = createUsersTable(dbPool); err != nil {
		log.Fatalf("Failed to create users table: %v\n", err)
	}

	channel <- dbPool
}
