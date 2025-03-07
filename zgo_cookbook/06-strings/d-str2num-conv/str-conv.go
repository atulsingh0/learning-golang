package main

import (
	"fmt"
	"strconv"
)

func main() {

	no, _ := strconv.Atoi("123")
	fmt.Println(no * 2)

	no2, _ := strconv.ParseInt("3849", 10, 64) // base 64 is for input
	fmt.Printf("%v is %T type\n", no2, no2)

	no, err := strconv.Atoi("Salting")
	if err != nil {
		e := err.(*strconv.NumError)
		fmt.Println(e.Err)
		fmt.Println(e.Func)
		fmt.Println(e.Num)
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(no)
	}

}
