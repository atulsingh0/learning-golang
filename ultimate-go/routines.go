package main

import (
	"fmt"
	"time"
)

func main() {

	defer timer()()
	for i := 0; i < 500; i++ {
		go increment(i)
	}
}

func increment(x int) {
	fmt.Println("now x is: ", x+1)
}

func timer() func() {
	fmt.Println("Timer started")
	start := time.Now()
	return func() {
		fmt.Println("Timer ended")
		fmt.Println("Time elapsed:", time.Since(start))
	}
}
