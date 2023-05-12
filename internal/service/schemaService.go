package service

import (
	"encoding/json"
	"io"
	"smartwayTestTAsk/internal/models"
	"smartwayTestTAsk/pkg/tools"
	"strconv"
)

func (s *service) AddSchema(rawData io.Reader) (err error) {
	newSchema := &models.Schema{}

	if err = tools.ParseJson(rawData, newSchema); err != nil {
		return
	}

	if err = s.repository.CreateSchema(newSchema); err != nil {
		return
	}
	return
}

func (s *service) FindSchemaByName(name string) ([]byte, error) {
	schema, err := s.repository.FindSchema(name)
	if err != nil {
		return nil, err
	}

	schemaProviders, err := s.repository.GetSchemaProviders(name)
	if err != nil {
		return nil, err
	}

	for _, value := range schemaProviders {
		schema.Providers = append(schema.Providers, value.ProviderId)
	}

	response, errMarshal := json.Marshal(schema)
	if errMarshal != nil {
		return nil, errMarshal
	}
	return response, nil
}

func (s *service) UpdateSchema(rawData io.Reader, id string) (err error) {
	schemaId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	newSchema := &models.Schema{}
	if err = tools.ParseJson(rawData, newSchema); err != nil {
		return
	}

	if newSchema.Name != "" {
		if err = s.repository.UpdateSchemaName(newSchema.Name, uint64(schemaId)); err != nil {
			return
		}
	}

	if newSchema.Providers != nil {
		if err = s.repository.DeleteFromSchemaProvider(uint64(schemaId)); err != nil {
			return
		}

		for _, value := range newSchema.Providers {
			providerId, errSelect := s.repository.SelectProviderId(value)
			if errSelect != nil {
				err = errSelect
				return
			}

			if err = s.repository.InsertToSchemaProvider(uint64(schemaId), providerId); err != nil {
				return
			}
		}
	}
	return
}

func (s *service) SaveDeleteSchema(id string) (err error) {
	schema_id, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	if err = s.repository.DeleteSchema(uint64(schema_id)); err != nil {
		return
	}
	return
}
