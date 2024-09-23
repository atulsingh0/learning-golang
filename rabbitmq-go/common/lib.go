package common

import (
	"bufio"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"math/rand"
	"os"
	"strings"
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	VHost    string `json:"vhost"`
	Queue    string `json:"queue"`
}

type RabbitMQ struct {
	Config Config
	Conn   *amqp.Connection
	Ch     *amqp.Channel
}

func NewRabbitMQ(config Config) *RabbitMQ {
	if config.Host == "" || config.Port == "" || config.User == "" || config.Password == "" || config.VHost == "" {
		log.Fatal("RabbitMQ config is not valid")
	}
	return &RabbitMQ{
		Config: config,
	}
}

func Connect(rq *RabbitMQ) error {
	var err error
	rq.Conn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:5672/%s",
		rq.Config.User, rq.Config.Password, rq.Config.Host, rq.Config.VHost))
	if err != nil {
		return err
	}
	fmt.Println("Connected to RabbitMQ. is connection closed ?: ", rq.Conn.IsClosed())
	return nil
}

func CreateChannel(rq *RabbitMQ) error {
	var err error

	if rq.Conn.IsClosed() {
		return err
	}

	rq.Ch, err = rq.Conn.Channel()
	if err != nil {
		return err
	}
	return nil
}

func CloseConnection(rq *RabbitMQ) {
	rq.Conn.Close()
}

func CloseChannel(rq *RabbitMQ) {
	rq.Ch.Close()
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Read_env_file() {
	env_file, err := os.Open(".env")
	if err != nil {
		log.Fatal(err)
	}
	defer env_file.Close()
	// Read the file line by line using scanner
	scanner := bufio.NewScanner(env_file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line by the = character
		parts := strings.Split(line, ": ")
		// Check if the line has the correct format
		if len(parts) == 2 {
			// Set the environment variable
			err = os.Setenv(parts[0], parts[1])
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
