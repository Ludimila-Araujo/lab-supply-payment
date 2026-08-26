package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Ludimila-Araujo/labsupply-payment/internal/service"
	"github.com/google/uuid"
)

type PaymentController struct {
	paymentService *service.PaymentService
}

func NewPaymentController(
	paymentService *service.PaymentService,
) *PaymentController {

	return &PaymentController{
		paymentService: paymentService,
	}
}

type CreatePaymentRequest struct {
	OrderID uuid.UUID `json:"order_id"`
	Amount  float64   `json:"amount"`
}

func (c *PaymentController) Create(
	w http.ResponseWriter,
	r *http.Request,
) {

	var request CreatePaymentRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	payment, err := c.paymentService.Create(
		request.OrderID,
		request.Amount,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(payment)
}

func (c *PaymentController) FindByID(
	w http.ResponseWriter,
	r *http.Request,
) {

	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid payment id", http.StatusBadRequest)
		return
	}

	payment, err := c.paymentService.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(payment)
}
