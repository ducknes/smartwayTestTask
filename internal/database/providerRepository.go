package database

import (
	"context"
	"fmt"
	"log"
	"smartwayTestTAsk/internal/models"

	"github.com/jackc/pgx/v5/pgconn"
)

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

func (repo *repository) SelectProviderId(provider_id string) (id uint64, err error) {
	if err = repo.database.QueryRow(context.Background(), SelectProviderId, provider_id).Scan(&id); err != nil {
		return 0, err
	}
	return
}

func (repo *repository) GetProviderAirlines(provider_id string) (airlineProviders []*models.AirlineProvider, err error) {
	rows, err := repo.database.Query(context.Background(), GetProviderAirlines, provider_id)
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
