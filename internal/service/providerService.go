package service

import (
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

func (s *service) GetAirlinesByProviderId(providerId int) {}
