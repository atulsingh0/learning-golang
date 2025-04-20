package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Reading from channel...")

	c := fanIn(someWork("AWS:"), someWork("GCP:"))
	// finite loop
	for i := 0; i < 20; i++ {
		fmt.Println(<-c)
	}
}

func fanIn(ch1, ch2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-ch1
		}
	}()

	go func() {
		for {
			c <- <-ch2
		}
	}()

	return c
}

// Generator: function that returns a channel
func someWork(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ { // infinite loop to send
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return ch
}
