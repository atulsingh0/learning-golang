package main

import (
	"fmt"
	"time"
)

// Controlling No of routines run at a time
func main() {

	counter := make(chan int, 4)
	start := time.Now()
	for i := 0; i < 40; i++ {
		counter <- 1
		j := i
		go func() {
			sleeping(1, j) // sleep 2 sec
			<-counter
		}()
	}
	fmt.Println("Total time:", time.Since(start))

	time.Sleep(time.Second * 15)
}

func sleeping(no int, count int) {
	time.Sleep(time.Duration(no) * time.Second)
	fmt.Printf("%v: Done\n", count)
}
