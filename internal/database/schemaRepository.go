package database

import (
	"context"
	"fmt"
	"log"
	"smartwayTestTAsk/internal/models"

	"github.com/jackc/pgx/v5/pgconn"
)

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

func (repo *repository) FindSchema(name string) (schema models.Schema, err error) {
	if err = repo.database.QueryRow(context.Background(), FindSchemaByName, name).Scan(&schema.Id, &schema.Name); err != nil {
		return
	}
	return
}

func (repo *repository) UpdateSchemaName(newName string, id uint64) error {
	if _, err := repo.database.Exec(context.Background(), UpdateSchemaName, newName, id); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			log.Println(newErr)
			return newErr
		}
		return err
	}
	return nil
}
