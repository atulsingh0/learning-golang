package main

import (
	"fmt"
	"reflect"
)

func main() {

	const (
		x = "yes"
		y // auto assume the value from previous def "yes"
		z = "No"
	)
	fmt.Println(x, y, z)
	fmt.Println("The type of 'x': ", reflect.TypeOf(x))

	var number = 5
	fmt.Println("number is : ", number, reflect.TypeOf(number))

	var data int8 = -124
	fmt.Println("data is : ", data, reflect.TypeOf(data))

	var udata uint8 = 124
	fmt.Println("udata is : ", udata, reflect.TypeOf(udata))

	enable := false
	fmt.Println("App is enabled:", enable, reflect.TypeOf(enable))

	const (
		no_1 int = 102 + iota // iota can be used with const type only
		no_2
		no_3
	)

	fmt.Println("iota magic:", no_1, no_2, no_3)
	fmt.Printf("iota magic: %T, %T, %T\n", no_1, no_2, no_3)

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
	)

	// golang dont raise concern abt const which is not used
	const pi = 3.14

	// default type is 0, "", false

	var (
		alpha int
		beta  float32
		gamma string
		delta bool
	)

	fmt.Println(alpha, beta, gamma, delta)

}
