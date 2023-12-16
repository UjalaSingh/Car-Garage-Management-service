package store

import (
	"sample/model"

	"gofr.dev/pkg/gofr"
)

type Car interface {
	// Create inserts a new car record into the database
	Create(ctx *gofr.Context, student *model.Car) (*model.Car, error)

	// GetByID retrieves a car record based on its ID
	GetByID(ctx *gofr.Context, id int) (*model.Car, error)

	// Update updates an existing car record with the provided information
	Update(ctx *gofr.Context, student *model.Car) (*model.Car, error)

	// Delete removes a car record from the database based on its ID
	Delete(ctx *gofr.Context, id int) error

	// GetAll retrieves all car records from the database
	GetAll(ctx *gofr.Context) ([]*model.Car, error)
}
