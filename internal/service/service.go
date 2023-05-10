package service

import (
	"io"
	"smartwayTestTAsk/internal/database"
)

type Service interface {
	AddAirline(rawData io.Reader) error
	DeleteAirlineByCode(iata string)
	AddProvider(rawData io.Reader) error
	DeleteProviderById(id string)
	UpdateProvidersList(iata string, providers []string)
	AddSchema(rawData io.Reader) error
	FindSchemaByName(name string)
	UpdateSchema(fields ...interface{})
	SaveDeleteSchema(id int)
	AddAccount(rawData io.Reader) error
	UptateAccountSchema(accountID int, schemaId int)
	DeleteAccountById(id int)
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
