package main

import (
	"log"
	"os"
	cmn "rmq-go/common"
	q "rmq-go/queue"
	"rmq-go/receiver"
)

var (
	vHost   string
	vUser   string
	vPass   string
	vCookie string
)

func init() {
	cmn.Read_env_file()
	vHost = os.Getenv("RABBITMQ_DEFAULT_VHOST")
	vUser = os.Getenv("RABBITMQ_DEFAULT_USER")
	vPass = os.Getenv("RABBITMQ_DEFAULT_PASS")
	vCookie = os.Getenv("RABBITMQ_ERL_COOKIE")
}

func main() {
	rmq := cmn.NewRabbitMQ(cmn.Config{
		Host:     "localhost",
		Port:     "5672",
		User:     vUser,
		Password: vPass,
		VHost:    vHost,
	})

	qname := "test-Queue"
	err := cmn.Connect(rmq)
	defer cmn.CloseConnection(rmq)
	cmn.FailOnError(err, "Failed to connect to RabbitMQ")

	err = cmn.CreateChannel(rmq)
	defer cmn.CloseChannel(rmq)
	cmn.FailOnError(err, "Failed to open a channel")

	// We create a Queue to send the message to.
	_, err = q.CreateQueueWithArgs(rmq, qname, nil)
	cmn.FailOnError(err, "Failed to declare a queue")

	// We set the payload for the message.
	ch, err := receiver.FetchMsg(rmq, qname)
	// If there is an error publishing the message, a log will be displayed in the terminal.
	cmn.FailOnError(err, "Failed to publish a message")
	for msg := range ch {
		log.Printf(" [x] Congrats, received message: %s", msg.Body)
	}

	//msg, err := rmq.fetchMsgQueue()
	//failOnError(err, "Failed to fetch message")
	//log.Printf(" [x] Congrats, received message: %s", msg)

}
