package events

import "github.com/google/uuid"

type PaymentApproved struct {
	OrderID   uuid.UUID `json:"order_id"`
	PaymentID uuid.UUID `json:"payment_id"`
}

type PaymentFailed struct {
	OrderID uuid.UUID `json:"order_id"`
	Reason  string    `json:"reason"`
}
