package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)
	ph := make(chan string)

	for i := 0; i < 10; i++ {
		j := i
		go func() { ch <- fmt.Sprintf("Routine1 %d", j) }()
		go func() { ph <- fmt.Sprintf("Routine2 %d", j*20) }()
	}

	// breaking the loop with timer
	timer := time.NewTimer(time.Second * 3)
outer:
	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		case msg := <-ph:
			fmt.Println(msg)
		case <-timer.C:
			fmt.Println("Timer expired")
			close(ch)
			close(ph)
			break outer
		}
	}
}
