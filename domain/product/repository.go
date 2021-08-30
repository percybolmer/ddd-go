// Package product holds the repository and the implementations for a ProductRepository
package product

import (
	"errors"

	"github.com/google/uuid"
	"github.com/percybolmer/ddd-go/aggregate"
)

var (
	//ErrProductNotFound is returned when a product is not found
	ErrProductNotFound = errors.New("the product was not found")
	//ErrProductAlreadyExist is returned when trying to add a product that already exists
	ErrProductAlreadyExist = errors.New("the product already exists")
)

// ProductRepository is the repository interface to fulfill to use the product aggregate
type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}
