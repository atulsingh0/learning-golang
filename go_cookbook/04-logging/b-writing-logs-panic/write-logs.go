package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	no := "78392340a" // change the value to see the difference
	out := convert(no)
	fmt.Println(out)
}

func convert(s string) int64 {
	defer log.Println("Defering from convert function")

	log.Println("Converting string to int")
	out, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		// raise panic
		log.Panicln("Can not parse string to int", err)
	}
	return out
}
