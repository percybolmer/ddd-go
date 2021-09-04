package tavern

import "github.com/google/uuid"

// Item represents a Item for all sub domains
type Item struct {
	id          uuid.UUID
	name        string
	description string
}
