package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)
	ph := make(chan string)

	// usind done channel
	done := make(chan bool)

	for i := 0; i < 12; i++ {
		j := i
		go func() { ch <- fmt.Sprintf("Routine1 %d", j) }()
		go func() { ph <- fmt.Sprintf("Routine2 %d", j*11) }()
	}

	go doWork(ch, ph, done)
	time.Sleep(2 * time.Second)

	close(done)
}

func doWork(ch, ph <-chan string, done <-chan bool) {
	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		case msg := <-ph:
			fmt.Println(msg)
		case <-done:
			return
		}
	}
}
