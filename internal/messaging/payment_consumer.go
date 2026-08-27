package messaging

import (
	"encoding/json"
	"log"

	"github.com/Ludimila-Araujo/labsupply-payment/internal/events"
	"github.com/Ludimila-Araujo/labsupply-payment/internal/service"

	"github.com/rabbitmq/amqp091-go"
)

type PaymentConsumer struct {
	paymentService *service.PaymentService
}

func NewPaymentConsumer(
	paymentService *service.PaymentService,
) *PaymentConsumer {

	return &PaymentConsumer{
		paymentService: paymentService,
	}
}

func (c *PaymentConsumer) Start(
	messages <-chan amqp091.Delivery,
) {

	for msg := range messages {

		var event events.PaymentRequested

		if err := json.Unmarshal(msg.Body, &event); err != nil {
			log.Printf(
				"invalid payment requested event: %v",
				err,
			)
			continue
		}

		_, err := c.paymentService.Create(
			event.OrderID,
			event.Amount,
		)

		if err != nil {
			log.Printf(
				"failed to create payment: %v",
				err,
			)
			continue
		}

		log.Printf(
			"payment created for order %s",
			event.OrderID,
		)
	}
}
