package notifier

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type Notifier struct {
	queueName  string
	connection *amqp.Channel
}

func NewNotifier(queueName string, connection *amqp.Channel) *Notifier {
	return &Notifier{
		queueName:  queueName,
		connection: connection,
	}
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
