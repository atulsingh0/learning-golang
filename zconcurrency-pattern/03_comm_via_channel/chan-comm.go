package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	ch := make(chan string)

	for i := 0; i < 10; i++ {
		j := i
		go someWork(j, ch)
		fmt.Println("Msg:", <-ch)
	}
	fmt.Println("I am done.")
	close(ch) // not required
	fmt.Println("-------------------")

	ph := make(chan string)
	go boringWork(99, ph)

	// looping only 15 times to fetch the data from channel
	for k := 0; k < 15; k++ {
		fmt.Printf("boringWork says: %s\n", <-ph)
	}
	close(ph)

}

func someWork(i int, ch chan string) {
	ch <- fmt.Sprintf("this is worker no: %d", i)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func boringWork(i int, ch chan string) {
	// infinite loop to send some msg to channel
	for j := 0; ; j++ {
		ch <- fmt.Sprintf("this is worker no: %d", j)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}
