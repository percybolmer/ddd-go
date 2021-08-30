// Package aggregate
// File: product.go
// Product is an aggregate that represents a product.
package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/percybolmer/ddd-go/entity"
)

var (
	// ErrMissingValues is returned when a product is created without a name or description
	ErrMissingValues = errors.New("missing values")
)

// Product is a aggregate that combines item with a price and quantity
type Product struct {
	// item is the root entity which is an item
	Item  *entity.Item `json:"item" bson:"item"`
	Price float64      `json:"price" bson:"price"`
	// Quantity is the number of products in stock
	Quantity int `json:"quantity" bson:"quantity"`
}

// NewProduct will create a new product
// will return error if name of description is empty
func NewProduct(name, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrMissingValues
	}

	return Product{
		Item: &entity.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		Price:    price,
		Quantity: 0,
	}, nil
}

func (p Product) GetID() uuid.UUID {
	return p.Item.ID
}

func (p Product) GetItem() *entity.Item {
	return p.Item
}
