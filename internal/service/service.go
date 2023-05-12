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
	UpdateProvidersList(data map[string]string) error
	AddSchema(rawData io.Reader) error
	FindSchemaByName(name string) ([]byte, error)
	UpdateSchema(rawData io.Reader, id string) error
	SaveDeleteSchema(id string) error
	AddAccount(rawData io.Reader) error
	UptateAccountSchema(data map[string]string) error
	DeleteAccountById(id string) error
	GetAirlinesByAccountId(id string) ([]byte, error)
	GetAirlinesByProviderId(id string) ([]byte, error)
}

type service struct {
	repository database.Storage
}

func NewService(repository database.Storage) Service {
	return &service{
		repository: repository,
	}
}
