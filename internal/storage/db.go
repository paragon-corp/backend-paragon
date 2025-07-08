// internal/storage/db.go
package storage

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ReadDB  *pgxpool.Pool
	WriteDB *pgxpool.Pool
)

func InitDB() {
	ctx := context.Background()

	writeDSN := os.Getenv("WRITE_DATABASE_URL")
	var err error
	WriteDB, err = pgxpool.New(ctx, writeDSN)
	if err != nil {
		log.Fatalf("❌ Unable to connect to write DB: %v", err)
	}
	log.Println("✅ Connected to WRITE DB")

	readDSN := os.Getenv("READ_DATABASE_URL")
	ReadDB, err = pgxpool.New(ctx, readDSN)
	if err != nil {
		log.Fatalf("❌ Unable to connect to read DB: %v", err)
	}
	log.Println("✅ Connected to READ DB")
}
