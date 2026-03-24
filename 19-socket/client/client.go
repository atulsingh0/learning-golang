package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Generating a string to write to connection
	msg := fmt.Sprintf("I am sending this msg to server at %s", time.Now())

	n, err := conn.Write([]byte(msg))
	if err != nil {
		return
	}
	log.Println("Sent Bytes: ", n)
}
