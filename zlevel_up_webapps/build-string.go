package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	forlioop()
	builder()
}

func forlioop() {
	defer timer()()

	var line string
	for i := 0; i < 200000; i++ {
		line += "hello"
	}
}

func builder() {
	defer timer()()

	var line strings.Builder
	for i := 0; i < 200000; i++ {
		line.WriteString("hello")
	}
	//fmt.Println(line.String())

}

func timer() func() {
	start := time.Now()
	return func() {
		fmt.Println("time:", time.Since(start))
	}
}
