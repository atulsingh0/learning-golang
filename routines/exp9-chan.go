package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	c := make(chan int)

	arr := randIntArr(100, 20)
	go squareSum(arr, c)
	sum := <-c
	fmt.Println(sum, arr)

	arr = randIntArr(100, 20)
	go squareSum(arr, c)
	sum = <-c
	fmt.Println(sum, arr)

	// close the channel
	// sender only should close the channel
	close(c)
}

func squareSum(arr []int, c chan int) {
	sum := 0
	for _, n := range arr {
		sum += n * n
	}
	// send the sum to the channel
	c <- sum
}

func randIntArr(m, n int) []int {
	return rand.Perm(m)[:n]
}
