package messaging

import (
	"encoding/json"
	"log"

	"github.com/Ludimila-Araujo/labsupply-payment/internal/events"
	"github.com/Ludimila-Araujo/labsupply-payment/internal/service"

	"github.com/rabbitmq/amqp091-go"
)

type PaymentConsumer struct {
	paymentService  *service.PaymentService
	resultPublisher PaymentResultPublisher
}

func NewPaymentConsumer(
	paymentService *service.PaymentService,
	resultPublisher PaymentResultPublisher,
) *PaymentConsumer {
	return &PaymentConsumer{
		paymentService:  paymentService,
		resultPublisher: resultPublisher,
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

		payment, err := c.paymentService.Create(
			event.OrderID,
			event.Amount,
		)

		if err != nil {
			log.Printf(
				"failed to create payment: %v",
				err,
			)

			if publishErr := c.resultPublisher.PublishPaymentFailed(
				events.PaymentFailed{
					OrderID: event.OrderID,
					Reason:  err.Error(),
				},
			); publishErr != nil {
				log.Printf(
					"failed to publish payment failed event: %v",
					publishErr,
				)
			}

			if err := msg.Ack(false); err != nil {
				log.Printf(
					"failed to acknowledge payment message: %v",
					err,
				)
			}

			continue
		}

		if err := c.resultPublisher.PublishPaymentApproved(
			events.PaymentApproved{
				OrderID:   event.OrderID,
				PaymentID: payment.ID,
			},
		); err != nil {
			log.Printf(
				"failed to publish payment approved event: %v",
				err,
			)

			continue
		}

		log.Printf(
			"payment created for order %s",
			event.OrderID,
		)

		if err := msg.Ack(false); err != nil {
			log.Printf(
				"failed to acknowledge payment message: %v",
				err,
			)
		}
	}
}
