package valueobject

import (
	"time"

	"github.com/google/uuid"
)

// Transaction is a payment between two parties
type Transaction struct {
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
