package tavern

import (
	"time"

	"github.com/google/uuid"
)

// Transaction is a payment between two parties
type Transaction struct {
	Amount    int       `json:"amount" bson:"amount"`
	From      uuid.UUID `json:"from" bson:"from"`
	To        uuid.UUID `json:"to" bson:"to"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}
