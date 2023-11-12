package configuration

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbitmq struct {
	queueName        string
	connectionString string
}

func NewRabbitmq(queueName, connectionString string) *Rabbitmq {
	return &Rabbitmq{
		connectionString: connectionString,
		queueName:        queueName,
	}
}

func (r *Rabbitmq) Connect() (*amqp.Connection, error) {
	conn, err := amqp.Dial(r.connectionString)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (r *Rabbitmq) CreateChannelAndQueue(connection *amqp.Connection) (*amqp.Channel, error) {
	ch, err := connection.Channel()

	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(r.queueName, false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return ch, nil
}
