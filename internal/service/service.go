package service

import (
	"io"
	"smartwayTestTAsk/internal/database"
)

type Service interface {
	AddAirline(rawData io.Reader) error
	DeleteAirlineByCode(iata string) error
	AddProvider(rawData io.Reader) error
	DeleteProviderById(id string) error
	UpdateProvidersList(iata string, providers []string)
	AddSchema(rawData io.Reader) error
	FindSchemaByName(name string)
	UpdateSchema(fields ...interface{})
	SaveDeleteSchema(id string) error
	AddAccount(rawData io.Reader) error
	UptateAccountSchema(accountID int, schemaId int)
	DeleteAccountById(id string) error
	GetAirlinesByAccountId(accountId int)
	GetAirlinesByProviderId(providerId int)
}

type service struct {
	repository database.Storage
}

func NewService(repository database.Storage) Service {
	return &service{
		repository: repository,
	}
}
