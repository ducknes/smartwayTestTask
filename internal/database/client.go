package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"smartwayTestTAsk/internal/config"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

type Client interface {
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, sc config.StorageConfig) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := goose.Up(db, config.MIGRPATH); err != nil {
		return nil, err
	}

	db.Close()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	pool, err = pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err != nil {
		log.Fatal("error do with tries postgresql")
	}

	return pool, err
}
