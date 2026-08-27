package messaging

import "github.com/rabbitmq/amqp091-go"

type RabbitMQ struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {

	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &RabbitMQ{
		conn:    conn,
		channel: channel,
	}, nil
}

func (r *RabbitMQ) Close() error {

	if err := r.channel.Close(); err != nil {
		r.conn.Close()
		return err
	}

	return r.conn.Close()
}

func (r *RabbitMQ) DeclarePaymentQueue() error {

	_, err := r.channel.QueueDeclare(
		"payments",
		true,
		false,
		false,
		false,
		nil,
	)

	return err
}

func (r *RabbitMQ) ConsumePayments() (<-chan amqp091.Delivery, error) {

	return r.channel.Consume(
		"payments",
		"",
		false,
		false,
		false,
		false,
		nil,
	)
}
