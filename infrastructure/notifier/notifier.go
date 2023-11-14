package notifier

import (
	"context"
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type Notifier struct {
	queueName  string
	connection *amqp.Channel
}

func NewNotifier(queueName string, connection *amqp.Channel) (*Notifier, error) {
	if queueName == "" {
		return &Notifier{}, errors.New("queueName must not be empty")
	}

	if connection == nil {
		return &Notifier{}, errors.New("connection must not be empty")
	}

	return &Notifier{
		queueName:  queueName,
		connection: connection,
	}, nil
}

func (n *Notifier) SendNotification(ctx context.Context, body string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := n.connection.PublishWithContext(ctx,
		"",
		n.queueName,
		false,
		false,
		amqp.Publishing{ContentType: "text/plain", Body: []byte(body)},
	)
	if err != nil {
		return err
	}

	log.Printf("message sent: %s\n", body)
	return nil
}
