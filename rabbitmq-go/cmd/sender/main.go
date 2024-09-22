package main

import (
	"log"
	"os"
	rm "rmq-go/common"
	q "rmq-go/queue"
	"rmq-go/sender"
	"strings"
)

var (
	vHost   string
	vUser   string
	vPass   string
	vCookie string
)

func init() {
	rm.Read_env_file()
	vHost = os.Getenv("RABBITMQ_DEFAULT_VHOST")
	vUser = os.Getenv("RABBITMQ_DEFAULT_USER")
	vPass = os.Getenv("RABBITMQ_DEFAULT_PASS")
	vCookie = os.Getenv("RABBITMQ_ERL_COOKIE")
}

func main() {
	rmq := rm.NewRabbitMQ(rm.Config{
		Host:     "localhost",
		Port:     "5672",
		User:     vUser,
		Password: vPass,
		VHost:    vHost,
	})

	qname := "test-Queue"
	err := rm.Connect(rmq)
	defer rm.CloseConnection(rmq)
	rm.FailOnError(err, "Failed to connect to RabbitMQ")

	err = rm.CreateChannel(rmq)
	defer rm.CloseChannel(rmq)
	rm.FailOnError(err, "Failed to open a channel")

	// We create a Queue to send the message to.
	_, err = q.CreateQueueWithArgs(rmq, qname, nil)
	rm.FailOnError(err, "Failed to declare a queue")

	// We set the payload for the message.
	// Reading the arguments from the command line
	body := strings.Join(os.Args[1:], " ")
	if body == "" {
		body = "Golang is awesome"
	}
	err = sender.Publish(rmq, qname, body)
	// If there is an error publishing the message, a log will be displayed in the terminal.
	rm.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Congrats, sending message: %s", body)

	//msg, err := rmq.fetchMsgQueue()
	//failOnError(err, "Failed to fetch message")
	//log.Printf(" [x] Congrats, received message: %s", msg)
}
