package main

import (
	"fmt"
	"strconv"
)

func main() {

	s := strconv.Itoa(2893)
	fmt.Println(s)

	s = strconv.FormatInt(392, 10)
	fmt.Println(s)

	no := 3.14159
	s = strconv.FormatFloat(no, 'f', 8, 64)
	fmt.Println(s)

}
