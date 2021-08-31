// Package product
// Product is an aggregate that represents a product.
package product

import (
	"errors"

	"github.com/google/uuid"
	"github.com/percybolmer/tavern"
)

var (
	// ErrMissingValues is returned when a product is created without a name or description
	ErrMissingValues = errors.New("missing values")
)

// Product is a aggregate that combines item with a price and quantity
type Product struct {
	// item is the root entity which is an item
	Item  *tavern.Item `json:"item" bson:"item"`
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
		Item: &tavern.Item{
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

func (p Product) GetItem() *tavern.Item {
	return p.Item
}
