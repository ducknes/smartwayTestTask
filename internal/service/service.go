package service

import (
	"smartwayTestTAsk/internal/database"
)

type Service interface {
	AddAirline()
	DeleteAirlineByCode(iata string)
	AddProvider()
	DeleteProviderById(id string)
	UpdateProvidersList(iata string, providers []string)
	AddSchema()
	FindSchemaByName(name string)
	UpdateSchema(fields ...interface{})
	SaveDeleteSchema(id int)
	AddAccount()
	UptateAccountSchema(accountID int, schemaId int)
	DeleteAccountById(id int)
	GetAirlinesByAccountId(accountId int)
	GetAirlinesByProviderId(providerId int)
}

type service struct {
	repository *database.Storage
}

func NewService(repository *database.Storage) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) AddAirline() {}

func (s *service) DeleteAirlineByCode(iata string) {}

func (s *service) AddProvider() {}

func (s *service) DeleteProviderById(id string) {}

func (s *service) UpdateProvidersList(iata string, providers []string) {}

func (s *service) AddSchema() {}

func (s *service) FindSchemaByName(name string) {}

func (s *service) UpdateSchema(fields ...interface{}) {}

func (s *service) SaveDeleteSchema(id int) {}

func (s *service) AddAccount() {}

func (s *service) UptateAccountSchema(accountID int, schemaId int) {}

func (s *service) DeleteAccountById(id int) {}

func (s *service) GetAirlinesByAccountId(accountId int) {}

func (s *service) GetAirlinesByProviderId(providerId int) {}
