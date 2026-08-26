package routes

import (
	"net/http"

	"github.com/Ludimila-Araujo/labsupply-payment/internal/controllers"
)

func Setup(
	paymentController *controllers.PaymentController,
) *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc(
		"POST /payments",
		paymentController.Create,
	)

	mux.HandleFunc(
		"GET /payments/{id}",
		paymentController.FindByID,
	)

	return mux
}
