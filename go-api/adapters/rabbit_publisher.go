package adapters

import (
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/eliastarin/ssp2-southpark/go-api/domain"
	"github.com/eliastarin/ssp2-southpark/go-api/ports"
)

type RabbitPublisher struct {
	conn  *amqp.Connection
	ch    *amqp.Channel
	queue string
}

var _ ports.MessagePublisher = (*RabbitPublisher)(nil)

func NewRabbitPublisher(amqpURL, queueName string) (*RabbitPublisher, error) {
	// 1) connect
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, err
	}

	// 2) channel
	ch, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		return nil, err
	}

	// 3) ensure queue exists (idempotent)
	_, err = ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // args
	)
	if err != nil {
		_ = ch.Close()
		_ = conn.Close()
		return nil, err
	}

	log.Printf("RabbitMQ: declared queue %s", queueName)
	return &RabbitPublisher{conn: conn, ch: ch, queue: queueName}, nil
}

func (p *RabbitPublisher) Publish(msg domain.Message) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return p.ch.PublishWithContext(ctx,
		"",      // default exchange
		p.queue, // routing key
		false,   // mandatory
		false,   // immediate (deprecated)
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         body,
		},
	)
}

func (p *RabbitPublisher) Close() {
	if p.ch != nil {
		_ = p.ch.Close()
	}
	if p.conn != nil {
		_ = p.conn.Close()
	}
}
