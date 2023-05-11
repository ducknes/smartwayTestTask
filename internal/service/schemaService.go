package service

import (
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

func (s *service) FindSchemaByName(name string) {}

func (s *service) UpdateSchema(fields ...interface{}) {}

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
