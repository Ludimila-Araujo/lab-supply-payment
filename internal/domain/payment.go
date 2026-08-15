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

	if amount <= 0 {
		return nil, ErrPaymentInvalidAmount
	}

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

//pedido aprovado:

func (p *Payment) Approve() error {

	if p.Status != PaymentStatusPending {
		return ErrPaymentInvalidStatus
	}

	p.Status = PaymentStatusApproved
	p.UpdatedAt = time.Now()

	return nil
}

//pedido negado:

func (p *Payment) Fail() error {

	if p.Status != PaymentStatusPending {
		return ErrPaymentInvalidStatus
	}

	p.Status = PaymentStatusFailed
	p.UpdatedAt = time.Now()

	return nil
}
