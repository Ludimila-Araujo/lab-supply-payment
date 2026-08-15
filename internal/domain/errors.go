package domain

import "errors"

var (
	ErrPaymentInvalidAmount = errors.New("payment amount must be greater than zero")
	ErrPaymentInvalidStatus = errors.New("payment status does not allow this operation")
)
