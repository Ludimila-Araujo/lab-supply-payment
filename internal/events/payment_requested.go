package events

import "github.com/google/uuid"

type PaymentRequested struct {
	OrderID uuid.UUID `json:"order_id"`
	Amount  float64   `json:"amount"`
}
