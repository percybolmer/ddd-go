package entity

import (
	"github.com/google/uuid"
)

// Person is a entity that represents a person in all Domains
type Person struct {
	// ID is the identifier of the Entity, the ID is shared for all sub domains
	ID uuid.UUID `json:"id" bson:"id"`
	// Name is the name of the person
	Name string `json:"name" bson:"name"`
	// Age is the age of the person
	Age int `json:"age" bson:"age"`
}
