package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)

	go func() {
		for _, d := range makeRange(1, 20) {
			ch <- d
		}
		defer close(ch)
	}()

	for d := range ch {
		fmt.Println(d)
	}

}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
