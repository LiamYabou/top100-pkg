package rabbitmq

import (
	"github.com/streadway/amqp"
)

func Open(amqpURL string) (amqpConn *amqp.Connection, err error) {
	amqpConn, err = amqp.Dial(amqpURL)
	return amqpConn, err
}
