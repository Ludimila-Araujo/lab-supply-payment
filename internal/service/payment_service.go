package service

import (
	"github.com/Ludimila-Araujo/labsupply-payment/internal/domain"
	"github.com/Ludimila-Araujo/labsupply-payment/internal/repository"
	"github.com/google/uuid"
)

type PaymentService struct {
	paymentRepository repository.PaymentRepository
}

func NewPaymentService(
	paymentRepository repository.PaymentRepository,
) *PaymentService {

	return &PaymentService{
		paymentRepository: paymentRepository,
	}
}

func (s *PaymentService) Create(
	orderID uuid.UUID,
	amount float64,
) (*domain.Payment, error) {

	payment, err := domain.NewPayment(orderID, amount)
	if err != nil {
		return nil, err
	}

	if err := s.paymentRepository.Create(payment); err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *PaymentService) FindByID(
	id uuid.UUID,
) (*domain.Payment, error) {

	return s.paymentRepository.FindByID(id)
}

func (s *PaymentService) Approve(
	id uuid.UUID,
) error {

	payment, err := s.paymentRepository.FindByID(id)
	if err != nil {
		return err
	}

	if err := payment.Approve(); err != nil {
		return err
	}

	return s.paymentRepository.Update(payment)
}

func (s *PaymentService) Fail(
	id uuid.UUID,
) error {

	payment, err := s.paymentRepository.FindByID(id)
	if err != nil {
		return err
	}

	if err := payment.Fail(); err != nil {
		return err
	}

	return s.paymentRepository.Update(payment)
}
