package database

import (
	"smartwayTestTAsk/internal/models"
)

type Storage interface {
	CreateAirline(airline *models.Airline) error
	CreateProvider(provider *models.Provider) error
	CreateSchema(schema *models.Schema) error
	CreateAccount(account *models.Account) error

	DeleteAirline(iata string) error
	DeleteProvider(provider_id string) error
	DeleteSchema(id uint64) error
	DeleteAccount(id uint64) error

	InsertToAirlineProvider(airline_id, provider_id uint64) error
	DeleteFromAirlineProvider(airline_id uint64) error
	SelectAirlineId(iata string) (uint64, error)
	SelectProviderId(provider_id string) (uint64, error)

	FindSchema(name string) (models.Schema, error)
	GetSchemaProviders(name string) ([]*models.SchemaProvider, error)
	UpdateSchemaName(newName string, id uint64) error
	DeleteFromSchemaProvider(schema_id uint64) error
	InsertToSchemaProvider(schema_id, provider_id uint64) error

	UpdateAccountSchema(schema_id, account_id uint64) error

	GetAccountAirlines(account_id uint64) ([]*models.AirlineProvider, error)
	GetProviderAirlines(provider_id string) ([]*models.AirlineProvider, error)
}
