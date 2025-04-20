package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	msg  string
	wait chan bool
}

// double check
// https://go.dev/talks/2012/concurrency.slide#30

func main() {

	fmt.Println("Reading from channel...")

	for i := 0; i < 8; i++ {
		ch := <-someWork("AWS", "CLI")
		ph := <-someWork("GCP", "GOLANG")
		fmt.Println(ch.msg)
		fmt.Println(ph.msg)
		ch.wait <- true
		ph.wait <- true
	}
}

// Generator: function that returns a channel
func someWork(name, msg string) <-chan Message {
	waitForIt := make(chan bool)
	ch := make(chan Message)
	go func() {
		for i := 0; ; i++ { // infinite loop to send
			ch <- Message{msg: fmt.Sprintf("%s %s: %d", name, msg, i), wait: waitForIt}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForIt
		}

	}()
	return ch
}
