package database

import (
	"context"
	"fmt"
	"log"
	"smartwayTestTAsk/internal/models"

	"github.com/jackc/pgx/v5/pgconn"
)

func (repo *repository) CreateAccount(account *models.Account) error {
	if _, err := repo.database.Exec(context.Background(), InsertAccount, account.SchemaId, account.Name); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			log.Println(newErr)
			return newErr
		}
		return err
	}
	return nil
}

func (repo *repository) DeleteAccount(id uint64) error {
	if _, err := repo.database.Exec(context.Background(), DeleteAccount, id); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			log.Println(newErr)
			return newErr
		}
		return err
	}
	return nil
}

func (repo *repository) UpdateAccountSchema(schema_id, account_id uint64) error {
	if _, err := repo.database.Exec(context.Background(), UpdateAccountSchema, schema_id, account_id); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			log.Println(newErr)
			return newErr
		}
		return err
	}
	return nil
}

func (repo *repository) GetAccountAirlines(account_id uint64) (airlineProviders []*models.AirlineProvider, err error) {
	rows, err := repo.database.Query(context.Background(), GetAccountAirlines, account_id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var airlineProvider models.AirlineProvider
		if err = rows.Scan(&airlineProvider.AirlineId, &airlineProvider.ProdviderId); err != nil {
			return
		}

		airlineProviders = append(airlineProviders, &airlineProvider)
	}

	return
}
