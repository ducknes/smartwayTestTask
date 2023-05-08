package database

import (
	"context"
	"fmt"
	"log"
	"smartwayTestTAsk/internal/models"

	"github.com/jackc/pgx/v5/pgconn"
)

type repository struct {
	client Client
}

func (repo *repository) Create(ctx context.Context, airline models.Airline) error {
	query := `INSERT INTO airline (iata, name) VALUES ($1, $2) RETURNING iata`

	if err := repo.client.QueryRow(ctx, query, airline.Iata, airline.Name).Scan(&airline.Iata); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			log.Println(newErr)
			return newErr
		}
		return err
	}
	return nil
}

//func (repo *repository) FindAll(ctx context.Context) (user []any, err error) {
//
//}
//
//func (repo *repository) FindOne(ctx context.Context, id string) (any, error) {
//
//}
//
//func (repo *repository) Update(ctx context.Context, user any) error {
//
//}
//
//func (repo *repository) Delete(ctx context.Context, id string) error {
//
//}

func NewRepository(client Client) Storage {
	return &repository{client: client}
}
