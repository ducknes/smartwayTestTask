package service

import (
	"io"
	"smartwayTestTAsk/internal/models"
	"smartwayTestTAsk/pkg/tools"
	"strconv"
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

func (s *service) DeleteAccountById(id string) (err error) {
	account_id, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	if err = s.repository.DeleteAccount(uint64(account_id)); err != nil {
		return
	}
	return
}

func (s *service) GetAirlinesByAccountId(accountId int) {}
