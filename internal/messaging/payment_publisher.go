package messaging

import (
	"encoding/json"

	"github.com/Ludimila-Araujo/labsupply-payment/internal/events"
	"github.com/rabbitmq/amqp091-go"
)

func (r *RabbitMQ) PublishPaymentRequested(
	event events.PaymentRequested,
) error {

	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return r.channel.Publish(
		"",
		"payments",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
