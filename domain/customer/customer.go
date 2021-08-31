// Package customer holds aggregates that combines many entities into a full object
package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/percybolmer/tavern"
)

var (
	// ErrInvalidPerson is returned when the person is not valid in the NewCustome factory
	ErrInvalidPerson = errors.New("a customer has to have an valid person")
)

// Customer is a aggregate that combines all entities needed to represent a customer
type Customer struct {
	// person is the root entity of a customer
	// which means the person.ID is the main identifier for this aggregate
	Person *tavern.Person `bson:"person"`
	// a customer can hold many products
	Products []*tavern.Item `bson:"products"`
	// a customer can perform many transactions
	Transactions []tavern.Transaction `bson:"transactions"`
}

// NewCustomer is a factory to create a new Customer aggregate
// It will validate that the name is not empty
func NewCustomer(name string) (Customer, error) {
	// Validate that the Name is not empty
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	// Create a new person and generate ID
	person := &tavern.Person{
		Name: name,
		ID:   uuid.New(),
	}
	// Create a customer object and initialize all the values to avoid nil pointer exceptions
	return Customer{
		Person:       person,
		Products:     make([]*tavern.Item, 0),
		Transactions: make([]tavern.Transaction, 0),
	}, nil
}

// GetID returns the customers root entity ID
func (c *Customer) GetID() uuid.UUID {
	return c.Person.ID
}

// SetName changes the name of the Customer
func (c *Customer) SetName(name string) {
	c.Person.Name = name
}
