package database

import (
	"context"
	"fmt"
	"log"
	"smartwayTestTAsk/internal/config"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Client interface {
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, sc config.StorageConfig) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)

	//err = DoWithTries(func() error {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	pool, err = pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	// 	return nil
	// }, sc.MaxAttempts, 1*time.Second)

	if err != nil {
		log.Fatal("error do with tries postgresql")
	}

	return pool, err
}

// func DoWithTries(fn func() error, attempts int, delay time.Duration) (err error) {
// 	for attempts > 0 {
// 		if err = fn(); err != nil {
// 			time.Sleep(delay)
// 			attempts--

// 			continue
// 		}
// 		return nil
// 	}
// 	return
// }
