package service

import (
	"io"
	"smartwayTestTAsk/internal/models"
	"smartwayTestTAsk/pkg/tools"
	"strings"
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

func (s *service) DeleteAirlineByCode(iata string) (err error) {
	if err = s.repository.DeleteAirline(iata); err != nil {
		return
	}
	return
}

func (s *service) UpdateProvidersList(data map[string]string) (err error) {
	airlineCode := data["airline_code"]
	newProvidersList := strings.Split(data["provider_ids"], ",")

	airline_id, errSelect := s.repository.SelectAirlineId(airlineCode)
	if errSelect != nil {
		err = errSelect
		return
	}

	if err = s.repository.DeleteFromAirlineProvider(airline_id); err != nil {
		return
	}

	for _, val := range newProvidersList {
		providerId, errSelect := s.repository.SelectProviderId(val)
		if errSelect != nil {
			err = errSelect
			return
		}

		if errInsert := s.repository.InsertToAirlineProvider(airline_id, providerId); errInsert != nil {
			err = errInsert
			return
		}
	}
	return
}
