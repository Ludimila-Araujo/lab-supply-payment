package service

import (
	"github.com/Ludimila-Araujo/labsupply-payment/internal/domain"
	"github.com/google/uuid"
)

type PaymentService struct{}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (s *PaymentService) Create(
	orderID uuid.UUID,
	amount float64,
) (*domain.Payment, error) {

	return domain.NewPayment(orderID, amount)
}

func (s *PaymentService) Approve(
	payment *domain.Payment,
) error {

	return payment.Approve()
}

func (s *PaymentService) Fail(
	payment *domain.Payment,
) error {

	return payment.Fail()
}
