package main

import "fmt"

const (
	noTickets int = 101
)

func main() {

	const (
		myConst       = 23
		noTickets int = 51
	)

	fmt.Printf("%v\n", myConst)

	//  Can not modify the value
	// noTickets = 23
	fmt.Printf("%v\n", noTickets) // Shodow the const noTickets

	fmt.Printf("%v, %T\n", myConst+noTickets, myConst+noTickets)

	var k int16 = 23
	fmt.Printf("%v, %T\n", myConst+k, myConst+k)
	// thow error mismatched types int and int16
	// fmt.Printf("%v, %T\n", noTickets+k, noTickets+k)

	//  declare by iota
	const x = iota // counter starting with 0
	fmt.Printf("%v, %T\n", x, x)

	const (
		y = iota
		z = iota
	)
	fmt.Printf("%v, %v\n", y, z)

	const (
		p = 5 + iota
		q
		r
	)
	fmt.Printf("%v, %v, %v\n", p, q, r)

	// Because iota start with value 0 which is the default value of int type var if not assigned
	// So avoid setting a constant value as 0

	var i int
	fmt.Println(i == y) // i = 0, and y = 0, hence true

}
