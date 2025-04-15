package main

import (
	"fmt"
	"time"
)

func main() {

	// Avoid this kind
	func() {
		start := time.Now()
		// below defer is calculated right away, avoid this kind if defer
		defer fmt.Println("Total time taken by onlyWait:", time.Since(start))
		onlyWait(5)
	}()

	// Fix the above but not safe
	func() {
		// use this kind
		start := time.Now()
		onlyWait(5)
		fmt.Println("Total time taken by onlyWait (fix):", time.Since(start))
	}()

	// Best one
	func() {
		waitFewSecond(5)
	}()

	time.Sleep(time.Second * 10)

}

func waitFewSecond(no int) {
	start := time.Now()
	time.Sleep(time.Duration(no) * time.Second)
	fmt.Println("Total time take by func (waitFewSecond):", time.Since(start))
}

func onlyWait(no int) {
	time.Sleep(time.Duration(no) * time.Second)
}
