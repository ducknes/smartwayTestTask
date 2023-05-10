package service

import (
	"io"
	"smartwayTestTAsk/internal/models"
	"smartwayTestTAsk/pkg/tools"
)

func (s *service) AddAirline(rawData io.Reader) (err error) {
	newAirline := &models.Airline{}

	if err = tools.ParseJson(rawData, newAirline); err != nil {
		return
	}

	if err = s.repository.CreateAirline(newAirline); err != nil {
		return
	}
	return
}

func (s *service) DeleteAirlineByCode(iata string) {}

func (s *service) UpdateProvidersList(iata string, providers []string) {}
