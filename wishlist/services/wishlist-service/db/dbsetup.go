package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func createTables(dbpool *pgxpool.Pool) error {
	query := `
	CREATE TABLE IF NOT EXISTS wishlists(
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		title VARCHAR(255) NOT NULL DEFAULT '',
		description VARCHAR(255) NOT NULL DEFAULT '',
		is_public boolean NOT NULL DEFAULT FALSE
	);
	CREATE TABLE IF NOT EXISTS items(
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		wishlist_id UUID,
		name VARCHAR(255) NOT NULL DEFAULT '',
		url TEXT NOT NULL DEFAULT '',
		price REAL NOT NULL DEFAULT 0,
		is_gifted boolean NOT NULL DEFAULT FALSE,
		gifted_by UUID,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(gifted_by) REFERENCES users(id) ON DELETE SET NULL,
		FOREIGN KEY(wishlist_id) REFERENCES wishlists(id) ON DELETE CASCADE
	);
	CREATE TABLE IF NOT EXISTS shared(
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		wishlist_id UUID NOT NULL,
		user_id UUID NOT NULL,
		is_owner BOOLEAN NOT NULL DEFAULT FALSE,
		can_edit boolean NOT NULL DEFAULT FALSE,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		last_opened TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		last_modified TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(wishlist_id) REFERENCES wishlists(id) ON DELETE CASCADE,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
	);
	CREATE INDEX IF NOT EXISTS idx_wishlists_id ON wishlists(id);
	CREATE INDEX IF NOT EXISTS idx_items_wishlist ON items (wishlist_id);
	CREATE INDEX IF NOT EXISTS idx_shared_shared_with ON shared (user_id);
	CREATE INDEX IF NOT EXISTS idx_shared_wishlist_id ON shared (wishlist_id);
	CREATE INDEX IF NOT EXISTS idx_shared_wishlist_id_shared_with ON shared(wishlist_id, user_id);
`
	_, err := dbpool.Exec(context.Background(), query)
	if err != nil {
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
	err = createTables(dbPool)
	if err != nil {
		log.Fatalf("Failed to create tables: %v\n", err)
	}

	channel <- dbPool
}
