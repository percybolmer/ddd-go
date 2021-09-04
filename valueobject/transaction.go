package valueobject

import (
	"time"

	"github.com/google/uuid"
)

// Transaction is a payment between two parties
type Transaction struct {
	Amount    int
	From      uuid.UUID
	To        uuid.UUID
	CreatedAt time.Time
}
