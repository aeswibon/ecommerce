package amqp

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// PaymentsInterface is a struct that holds the connection, queue, channel and service
type PaymentsInterface struct {
	conn    *amqp.Connection
	queue   amqp.Queue
	channel *amqp.Channel

	// service application.PaymentsService
}

// NewPaymentsInterface is a constructor for PaymentsInterface
func NewPaymentsInterface(url, queueName string) (PaymentsInterface, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return PaymentsInterface{}, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return PaymentsInterface{}, err
	}
	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return PaymentsInterface{}, err
	}
	return PaymentsInterface{conn, q, ch}, nil
}

// Run is a method that runs the service
func (p *PaymentsInterface) Run(ctx context.Context) error {
	msgs, err := p.channel.Consume(
		p.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	defer p.Close()
	for {
		select {
		case <-ctx.Done():
			return nil
		case msg := <-msgs:
			// p.service.ProcessOrder(msg.Body)
			if err := msg.Ack(false); err != nil {
				log.Println("Error acknowledging message: ", err)
			}
		}
	}
}

// Close is a method that closes the connection
func (p *PaymentsInterface) Close() {
	if err := p.channel.Close(); err != nil {
		log.Println("Error closing channel: ", err)
	}
	if err := p.conn.Close(); err != nil {
		log.Println("Error closing connection: ", err)
	}
}
