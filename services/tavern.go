package services

import (
	"log"

	"github.com/google/uuid"
)

// TavernConfiguration is an alias that takes a pointer and modifies the Tavern
type TavernConfiguration func(os *Tavern) error

type Tavern struct {
	// orderservice is used to handle orders
	OrderService *OrderService
	// BillingService is used to handle billing
	// This is up to you to implement
	BillingService interface{}
}

// NewTavern takes a variable amount of TavernConfigurations and builds a Tavern
func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error) {
	// Create the Tavern
	t := &Tavern{}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

// WithOrderService applies a given OrderService to the Tavern
func WithOrderService(os *OrderService) TavernConfiguration {
	// return a function that matches the TavernConfiguration signature
	return func(t *Tavern) error {
		t.OrderService = os
		return nil
	}
}

// Order performs an order for a customer
func (t *Tavern) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}
	log.Printf("Bill the Customer: %0.0f", price)

	// Bill the customer
	//err = t.BillingService.Bill(customer, price)
	return nil
}
