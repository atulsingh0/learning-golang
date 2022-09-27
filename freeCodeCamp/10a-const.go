package main

import (
	"fmt"
	"reflect"
)

func main() {

	const tickets = 50

	const (
		i = 40
		j = 30
		k = 20
	)

	const (
		info = iota
		warn
		err
	)

	const (
		roll = 1 + iota
		roll2
		roll3
	)

	const (
		active  = 1 << iota // bitwise operation 1 << 0 = 1 => 1
		send                // 1 << 1 = 10 => 2
		receive             // 1 << 2 = 100 => 4
	)

	const (
		// value 9 will be set to all
		red = 9
		blue
		green
	)

	const (
		x = 5
		y = 3
		z
	)

	fmt.Println(tickets, reflect.TypeOf(tickets))
	fmt.Println(i, reflect.TypeOf(i))
	fmt.Println(j, reflect.TypeOf(j))
	fmt.Println(k, reflect.TypeOf(k))
	fmt.Println(info, reflect.TypeOf(info))
	fmt.Println(warn, reflect.TypeOf(warn))
	fmt.Println(err, reflect.TypeOf(err))

	fmt.Println(roll)
	fmt.Println(roll2)
	fmt.Println(roll3)

	fmt.Println(active)
	fmt.Println(send)
	fmt.Println(receive)

	fmt.Println(red)
	fmt.Println(blue)
	fmt.Println(green)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
