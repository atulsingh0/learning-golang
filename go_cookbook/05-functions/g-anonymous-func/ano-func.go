package main

import (
	"fmt"
	"reflect"
)

type magic func(int) int // magic is a function type which takes an int and returns an int

type data struct {
	no  string
	out magic
}

func main() {

	// anonymous function
	sq := func(x int) int {
		return x * x
	}

	fmt.Println(sq(10))
	fmt.Println(reflect.TypeOf(sq))
	fmt.Println(reflect.TypeOf(sq).Kind())

	test := data{
		no: "test",
		out: func(x int) int {
			return x * x
		},
	}

	fmt.Println(test.out(5))

}
