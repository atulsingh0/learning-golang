package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	no := "78392340" // change the value to see the difference

	out, err := strconv.ParseInt(no, 10, 64)
	if err != nil {
		log.Fatalln("Can not parse string to int", err)
	}
	fmt.Println(out)
}
