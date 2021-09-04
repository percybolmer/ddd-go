package entity

import "github.com/google/uuid"

// Item represents a Item for all sub domains
type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
}
