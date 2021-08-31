package tavern

import "github.com/google/uuid"

// Item represents a Item for all sub domains
type Item struct {
	ID          uuid.UUID `json:"id" bson:"id"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
}
