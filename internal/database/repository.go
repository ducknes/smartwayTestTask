package database

import (
	"context"
	"fmt"
	"log"
	"smartwayTestTAsk/internal/models"

	"github.com/jackc/pgx/v5/pgconn"
)

type repository struct {
	database Client
}

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

func (repo *repository) CreateProvider(provider *models.Provider) error {
	if _, err := repo.database.Exec(context.Background(), InsertProvider, provider.ProdviderId, provider.Name); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			log.Println(newErr)
			return newErr
		}
		return err
	}
	return nil
}

func (repo *repository) CreateSchema(schema *models.Schema) error {
	if _, err := repo.database.Exec(context.Background(), InsertSchema, schema.Name); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			log.Println(newErr)
			return newErr
		}
		return err
	}
	return nil
}

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

func (repo *repository) DeleteProvider(provider_id string) error {
	if _, err := repo.database.Exec(context.Background(), DeleteProvider, provider_id); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			log.Println(newErr)
			return newErr
		}
		return err
	}
	return nil
}

func (repo *repository) DeleteSchema(id uint64) error {
	if _, err := repo.database.Exec(context.Background(), DeleteSchema, id); err != nil {
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

func NewRepository(client Client) Storage {
	return &repository{database: client}
}
