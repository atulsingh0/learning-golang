package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Reading from channel...")

	ch := someWork("I am goroutine")

	// finite loop
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
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
