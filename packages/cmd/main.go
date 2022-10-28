package main

import (
	"fmt"

	"packages/services"
)

func init() {
	fmt.Println("Calling Init in Main")
}

func main() {

	fmt.Println("Square of 4 is:", services.Square(4))
	fmt.Println("String version of [1,2,4] is:", services.IntArrToStrArr([]int{1, 2, 4}))
}
