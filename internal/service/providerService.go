package service

import (
	"encoding/json"
	"io"
	"smartwayTestTAsk/internal/models"
	"smartwayTestTAsk/pkg/tools"
)

func (s *service) AddProvider(rawData io.Reader) (err error) {
	newProvider := &models.Provider{}

	if err = tools.ParseJson(rawData, newProvider); err != nil {
		return
	}

	if err = s.repository.CreateProvider(newProvider); err != nil {
		return
	}
	return
}

func (s *service) DeleteProviderById(id string) (err error) {
	if err = s.repository.DeleteProvider(id); err != nil {
		return
	}
	return
}

func (s *service) GetAirlinesByProviderId(id string) ([]byte, error) {
	airlineProviders, err := s.repository.GetProviderAirlines(id)
	if err != nil {
		return nil, err
	}

	response := make(map[string][]string, 1)
	airlines := make([]string, 0)

	for _, values := range airlineProviders {
		airlines = append(airlines, values.AirlineId)
	}

	response[id] = airlines

	jsonB, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	return jsonB, nil
}
