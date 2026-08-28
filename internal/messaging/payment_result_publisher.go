package messaging

import (
	"encoding/json"

	"github.com/Ludimila-Araujo/labsupply-payment/internal/events"
	"github.com/rabbitmq/amqp091-go"
)

type PaymentResultPublisher interface {
	PublishPaymentApproved(event events.PaymentApproved) error
	PublishPaymentFailed(event events.PaymentFailed) error
}

func (r *RabbitMQ) PublishPaymentApproved(
	event events.PaymentApproved,
) error {

	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return r.channel.Publish(
		"",
		"payment.approved",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

func (r *RabbitMQ) PublishPaymentFailed(
	event events.PaymentFailed,
) error {

	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return r.channel.Publish(
		"",
		"payment.failed",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
