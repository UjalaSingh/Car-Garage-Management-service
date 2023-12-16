package store

import (
	"database/sql"
	"fmt"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"sample/model"
)

type car struct{}

func New() *car {
	return &car{}
}

// Create inserts a new car record into the database
func (s *car) Create(ctx *gofr.Context, car *model.Car) (*model.Car, error) {
	_, err := ctx.DB().ExecContext(ctx, createQuery, car.ID, car.Name, car.Color)
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	return car, nil
}

// GetByID retrieves a car record based on its ID
func (s *car) GetByID(ctx *gofr.Context, id int) (*model.Car, error) {
	var resp model.Car

	err := ctx.DB().QueryRowContext(ctx, getByIDQuery, id).
		Scan(&resp.ID, &resp.Name, &resp.Color)
	switch err {
	case sql.ErrNoRows:
		return nil, errors.EntityNotFound{Entity: "car", ID: fmt.Sprintf("%v", id)}
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

// Update updates an existing car record with the provided information
func (s *car) Update(ctx *gofr.Context, car *model.Car) (*model.Car, error) {
	_, err := ctx.DB().ExecContext(ctx, updateQuery, car.Name, car.Color, car.ID)
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	return car, nil
}

// Delete removes a car record from the database based on its ID
func (s *car) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, deleteQuery, id)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}

// GetAll retrieves all car records from the database
func (s *car) GetAll(ctx *gofr.Context) ([]*model.Car, error) {
	rows, err := ctx.DB().QueryContext(ctx, getAllQuery)
	if err != nil {
		return nil, errors.DB{Err: err}
	}
	defer rows.Close()

	var resp []*model.Car
	for rows.Next() {
		var car model.Car
		if err := rows.Scan(&car.ID, &car.Name, &car.Color); err != nil {
			return nil, errors.DB{Err: err}
		}

		resp = append(resp, &car)
	}

	return resp, nil
}
