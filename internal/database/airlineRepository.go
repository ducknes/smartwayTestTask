package database

import (
	"context"
	"fmt"
	"log"
	"smartwayTestTAsk/internal/models"

	"github.com/jackc/pgx/v5/pgconn"
)

func (repo *repository) CreateAirline(airline *models.Airline) error {
	if _, err := repo.database.Exec(context.Background(), InsertAirline, airline.Iata, airline.Name); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			log.Println(newErr)
			return newErr
		}
		return err
	}
	return nil
}

func (repo *repository) DeleteAirline(iata string) error {
	if _, err := repo.database.Exec(context.Background(), DeleteAirline, iata); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			log.Println(newErr)
			return newErr
		}
		return err
	}
	return nil
}

func (repo *repository) SelectAirlineId(iata string) (airline_id uint64, err error) {
	if err = repo.database.QueryRow(context.Background(), SelectAirlineId, iata).Scan(&airline_id); err != nil {
		return 0, err
	}
	return
}
