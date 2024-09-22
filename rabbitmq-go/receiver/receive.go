package receiver

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"rmq-go/common"
)

func FetchMsg(rq *common.RabbitMQ, queue string) (<-chan amqp.Delivery, error) {
	var err error
	if rq.Conn.IsClosed() {
		return nil, err
	}
	msgs, err := rq.Ch.Consume(
		queue, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	return msgs, err
}
