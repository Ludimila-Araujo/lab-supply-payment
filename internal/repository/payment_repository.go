package repository

import (
	"github.com/Ludimila-Araujo/labsupply-payment/internal/domain"
	"github.com/google/uuid"
)

type PaymentRepository interface {
	Create(payment *domain.Payment) error
	FindByID(id uuid.UUID) (*domain.Payment, error)
	Update(payment *domain.Payment) error
}
