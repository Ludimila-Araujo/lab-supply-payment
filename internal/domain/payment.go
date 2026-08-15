package domain

import (
	"time"

	"github.com/google/uuid"
)

type PaymentStatus string

const (
	PaymentStatusPending  PaymentStatus = "PENDING"
	PaymentStatusApproved PaymentStatus = "APPROVED"
	PaymentStatusFailed   PaymentStatus = "FAILED"
)

type Payment struct {
	ID        uuid.UUID
	OrderID   uuid.UUID
	Amount    float64
	Status    PaymentStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPayment(orderID uuid.UUID, amount float64) (*Payment, error) {

	now := time.Now()

	return &Payment{
		ID:        uuid.New(),
		OrderID:   orderID,
		Amount:    amount,
		Status:    PaymentStatusPending,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
