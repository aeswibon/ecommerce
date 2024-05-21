package payments

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// AMQPSvc is a struct that holds the queue and channel
type AMQPSvc struct {
	queue amqp.Queue
	channel *amqp.Channel
}

// NewAMQPSvc is a constructor for AMQPSvc
func NewAMQPSvc(url, queuName string) (AMQPSvc, error) {
	conn, err := amqp.Dial(url); if err != nil {
		return AMQPSvc{}, err
	}
	ch, err := conn.Channel(); if err != nil {
		return AMQPSvc{}, err
	}
	q, err := ch.QueueDeclare(
		queuName,
		true,
		false,
		false,
		false,
		nil,
	); if err != nil {
		return AMQPSvc{}, err
	}
	return AMQPSvc{q, ch}, nil
}

// InitializeOrderPayments initializes the order payments
func (a *AMQPSvc) InitializeOrderPayments (id string, price float32) error {
	err := a.channel.Publish(
		"",
		a.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte("Order ID: " + id + " Price: " + string(rune(price))),
		},
	)
	if err != nil {
		log.Fatalf("Error publishing message: %s", err)
	}
	log.Printf("sent order %s to amqp", id)
	return nil
}

// Close closes the channel
func (a *AMQPSvc) Close () error {
	return a.channel.Close()
}