package service

import (
	"io"
	"smartwayTestTAsk/internal/models"
	"smartwayTestTAsk/pkg/tools"
)

func (s *service) AddAccount(rawData io.Reader) (err error) {
	newAccount := &models.Account{}

	if err = tools.ParseJson(rawData, newAccount); err != nil {
		return
	}

	if err = s.repository.CreateAccount(newAccount); err != nil {
		return
	}
	return
}

func (s *service) UptateAccountSchema(accountID int, schemaId int) {}

func (s *service) DeleteAccountById(id int) {}

func (s *service) GetAirlinesByAccountId(accountId int) {}
