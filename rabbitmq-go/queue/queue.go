package queue

import (
	"github.com/rabbitmq/amqp091-go"
	"rmq-go/common"
)

func CreateQueueWithArgs(rq *common.RabbitMQ, name string, args amqp091.Table) (amqp091.Queue, error) {
	var err error
	q, err := rq.Ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		args,  // arguments
	)
	if err != nil {
		return amqp091.Queue{}, err
	}
	return q, nil
}

func DeleteQueue(rq *common.RabbitMQ, name string) error {
	var err error
	_, err = rq.Ch.QueueDelete(
		name,  // name
		false, // ifUnused
		false, // ifEmpty
		false, // noWait
	)
	if err != nil {
		return err
	}
	return nil
}
