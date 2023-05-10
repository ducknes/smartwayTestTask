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
}
