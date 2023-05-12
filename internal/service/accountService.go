package service

import (
	"encoding/json"
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

func (s *service) UptateAccountSchema(data map[string]string) error {
	accountId, err := strconv.Atoi(data["account_id"])
	if err != nil {
		return err
	}

	schemaId, err := strconv.Atoi(data["schema_id"])
	if err != nil {
		return err
	}

	if err = s.repository.UpdateAccountSchema(uint64(schemaId), uint64(accountId)); err != nil {
		return err
	}
	return nil
}

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

func (s *service) GetAirlinesByAccountId(id string) ([]byte, error) {
	accountId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	airlineProviders, err := s.repository.GetAccountAirlines(uint64(accountId))
	if err != nil {
		return nil, err
	}

	response := make(map[uint64][]string, 1)
	airlines := make([]string, 0)

	uniqueFilter := make(map[string]bool)
	for _, value := range airlineProviders {
		if !uniqueFilter[value.AirlineId] {
			uniqueFilter[value.AirlineId] = true
			airlines = append(airlines, value.AirlineId)
		}
	}

	response[uint64(accountId)] = airlines

	jsonB, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return jsonB, nil
}
