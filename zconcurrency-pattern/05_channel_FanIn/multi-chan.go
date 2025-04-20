package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Reading from channel...")

	ch := someWork("AWS:")
	ph := someWork("GCP:")

	// finite loop
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch) // this is a blocking operation for ph
		fmt.Println(<-ph)
	}
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
