package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("wow: ", reflect.TypeOf("wow"))

	fmt.Println(34, reflect.TypeOf(34))
	fmt.Println(3.14, reflect.TypeOf(3.14))
	fmt.Println('A', reflect.TypeOf('A'))
	fmt.Println(true, reflect.TypeOf(true))

}
