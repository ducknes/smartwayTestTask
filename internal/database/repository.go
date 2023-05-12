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

func NewRepository(client Client) Storage {
	return &repository{database: client}
}

func (repo *repository) InsertToAirlineProvider(airline_id, provider_id uint64) error {
	if _, err := repo.database.Exec(context.Background(), InsertToAirlineProvires, airline_id, provider_id); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			log.Println(newErr)
			return newErr
		}
		return err
	}
	return nil
}

func (repo *repository) DeleteFromAirlineProvider(airline_id uint64) error {
	if _, err := repo.database.Exec(context.Background(), DeleteFromAirlineProvider, airline_id); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			log.Println(newErr)
			return newErr
		}
		return err
	}
	return nil
}

func (repo *repository) GetSchemaProviders(name string) ([]*models.SchemaProvider, error) {
	rows, err := repo.database.Query(context.Background(), GetSchemaProviders, name)
	if err != nil {
		return nil, err
	}

	var schemaProviders []*models.SchemaProvider
	for rows.Next() {
		var schemaProvider models.SchemaProvider
		err = rows.Scan(&schemaProvider.SchemaId, &schemaProvider.ProviderId)
		if err != nil {
			return nil, err
		}

		schemaProviders = append(schemaProviders, &schemaProvider)
	}
	rows.Close()

	return schemaProviders, nil
}

func (repo *repository) DeleteFromSchemaProvider(schema_id uint64) error {
	if _, err := repo.database.Exec(context.Background(), DeleteFromSchemaProvider, schema_id); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			log.Println(newErr)
			return newErr
		}
		return err
	}
	return nil
}

func (repo *repository) InsertToSchemaProvider(schema_id, provider_id uint64) error {
	if _, err := repo.database.Exec(context.Background(), InsertToSchemaProvires, provider_id, schema_id); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			log.Println(newErr)
			return newErr
		}
		return err
	}
	return nil
}
