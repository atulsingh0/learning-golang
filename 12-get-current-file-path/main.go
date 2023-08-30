package main

import (
	"fmt"
	"runtime"
)

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		fmt.Println("Caller:", filename)
	} else {
		fmt.Println("No idea what to do")
	}
}
