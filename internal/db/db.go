package db

import (
	db "7h3-3mp7y-m4n/INote/internal/db/generated"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool    *pgxpool.Pool
	Queries *db.Queries
}

func NewDB() (*DB, error) {
	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:5432/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_DB"))

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, fmt.Errorf("cannot parse config: %w", err)
	}

	config.MaxConns = 10
	config.MaxConnLifetime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("cannot create connection pool: %w", err)
	}

	return &DB{
		Pool:    pool,
		Queries: db.New(pool),
	}, nil
}
