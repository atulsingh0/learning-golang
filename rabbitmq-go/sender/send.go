package sender

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"rmq-go/common"
)

func Publish(rq *common.RabbitMQ, queue, msg string) error {
	var err error
	if rq.Conn.IsClosed() {
		return err
	}
	return rq.Ch.Publish(
		"",    // exchange
		queue, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
}
