package database

import (
	"context"
	"smartwayTestTAsk/internal/models"
)

// TODO: add all methods
type Storage interface {
	Create(ctx context.Context, airline models.Airline) error
	// FindAll(ctx context.Context) (user []any, err error)
	// FindOne(ctx context.Context, id string) (any, error)
	// Update(ctx context.Context, user any) error
	// Delete(ctx context.Context, id string) error
}
