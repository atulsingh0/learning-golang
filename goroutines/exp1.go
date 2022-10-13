package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		go func() {
			x := i
			time.Sleep(time.Second * 1)
			fmt.Println("Hi", i, x)
		}()
	}

	time.Sleep(time.Second * 2)

}
