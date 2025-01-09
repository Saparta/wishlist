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
		user_id UUID,
		title VARCHAR(255) NOT NULL,
		description VARCHAR(255),
		is_public boolean NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		last_modified TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		last_opened TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
	);
	CREATE TABLE IF NOT EXISTS items(
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		wishlist_id UUID,
		name VARCHAR(255) NOT NULL,
		url TEXT,
		price REAL,
		is_gifted boolean NOT NULL,
		gifted_by UUID,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(gifted_by) REFERENCES users(id) ON DELETE SET NULL,
		FOREIGN KEY(wishlist_id) REFERENCES wishlists(id) ON DELETE CASCADE
	);
	CREATE TABLE IF NOT EXISTS shared(
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		wishlist_id UUID NOT NULL,
		shared_with UUID NOT NULL,
		can_edit boolean NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		last_opened TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		last_modified TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(wishlist_id) REFERENCES wishlists(id) ON DELETE CASCADE,
		FOREIGN KEY(shared_with) REFERENCES users(id) ON DELETE CASCADE
	);
	CREATE INDEX IF NOT EXISTS idx_wishlist_user ON wishlists (id, user_id);
	CREATE INDEX IF NOT EXISTS idx_items_wishlist ON items (wishlist_id);
	CREATE INDEX IF NOT EXISTS idx_shared_shared_with ON shared (shared_with);
	CREATE INDEX IF NOT EXISTS idx_shared_wishlist_id ON shared (wishlist_id);
	CREATE OR REPLACE FUNCTION update_wishlist_last_modified()
	RETURNS TRIGGER AS $$
	BEGIN
    UPDATE wishlists
    SET last_modified = CURRENT_TIMESTAMP
    WHERE id = NEW.wishlist_id;
    
    RETURN NEW;
	END;
	$$ LANGUAGE plpgsql;
	CREATE OR REPLACE FUNCTION update_wishlist_last_modified_self()
	RETURNS TRIGGER AS $$
	BEGIN
    NEW.last_modified := CURRENT_TIMESTAMP;
    RETURN NEW; -- Return the updated row to finalize the change
	END;
	$$ LANGUAGE plpgsql;
	DO $$
	BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_trigger
        WHERE tgname = 'trigger_update_wishlist_last_modified'
    ) THEN
        CREATE TRIGGER trigger_update_wishlist_last_modified
        AFTER UPDATE ON items
        FOR EACH ROW
        WHEN (OLD.* IS DISTINCT FROM NEW.*)
        EXECUTE FUNCTION update_wishlist_last_modified();
    END IF;
    IF NOT EXISTS (
        SELECT 1
        FROM pg_trigger
        WHERE tgname = 'trigger_update_last_modified_self'
    ) THEN
        CREATE TRIGGER trigger_update_last_modified_self
        BEFORE UPDATE ON wishlists
        FOR EACH ROW
        WHEN (
            OLD.title IS DISTINCT FROM NEW.title OR
            OLD.description IS DISTINCT FROM NEW.description OR
            OLD.is_public IS DISTINCT FROM NEW.is_public
        )
        EXECUTE FUNCTION update_wishlist_last_modified_self();
    END IF;
END;
$$;
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
