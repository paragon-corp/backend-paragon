package storage

import (
	"context"
	"log"
	"time"
)

func InitSchema() {
	schema := `
CREATE TABLE IF NOT EXISTS coffees (
	id TEXT PRIMARY KEY,
	name TEXT NOT NULL,
	price NUMERIC NOT NULL
);

CREATE TABLE IF NOT EXISTS orders (
	id TEXT PRIMARY KEY,
	coffee_id TEXT NOT NULL REFERENCES coffees(id),
	status TEXT NOT NULL
);
`

	log.Println("Running schema migration...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := WriteDB.Exec(ctx, schema)
	if err != nil {
		log.Fatalf("❌ Failed to initialize schema: %v", err)
	}

	log.Println("✅ Schema initialized")
}
