package main

import (
	"fmt"
	"os"
	rm "rmq-go/common"
	q "rmq-go/queue"
	"rmq-go/sender"
	"strconv"
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
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		printUsage()
		os.Exit(1)
	}

	rmq := rm.NewRabbitMQ(rm.Config{
		Host:     "localhost",
		Port:     "5672",
		User:     vUser,
		Password: vPass,
		VHost:    vHost,
	})

	qname := "test-Queue"
	err = rm.Connect(rmq)
	defer rm.CloseConnection(rmq)
	rm.FailOnError(err, "Failed to connect to RabbitMQ")

	err = rm.CreateChannel(rmq)
	defer rm.CloseChannel(rmq)
	rm.FailOnError(err, "Failed to open a channel")

	// We create a Queue to send the message to.
	_, err = q.CreateQueueWithArgs(rmq, qname, nil)
	rm.FailOnError(err, "Failed to declare a queue")

	// We set the payload for the message.
	// sending the random generated message
	fmt.Println("input msgs: ", n)
	for i := 0; i < n; i++ {
		j := i
		if j > 20 {
			j = i % 20
		}
		j++
		err = sender.Publish(rmq, qname, rm.RandomString(j))
		//rm.FailOnError(err, "Failed to publish a message: "+inpts[i])
		//log.Printf(" [x] Congrats, sending message: %s", inpts[i])
	}
	// If there is an error publishing the message, a log will be displayed in the terminal.

	//msg, err := rmq.fetchMsgQueu
	//e()
	//failOnError(err, "Failed to fetch message")
	//log.Printf(" [x] Congrats, received message: %s", msg)
}

func printUsage() {
	fmt.Println("Usage: go run main.go <number of messages>")
}
