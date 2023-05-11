package database

import (
	"smartwayTestTAsk/internal/models"
)

// TODO: add all methods
type Storage interface {
	CreateAirline(airline *models.Airline) error
	CreateProvider(provider *models.Provider) error
	CreateSchema(schema *models.Schema) error
	CreateAccount(account *models.Account) error
	
	DeleteAirline(iata string) error
	DeleteProvider(provider_id string) error
	DeleteSchema(id uint64) error
	DeleteAccount(id uint64) error
}
