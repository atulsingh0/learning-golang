package main

import "fmt"

type point struct {
	x, y int
}

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	p := point{3, 4} // initializing custom type var

	fmt.Printf("%v\n", p)

	me := person{"atul", "singh", 30}

	fmt.Printf("%v\t%T\n", me, me)

	fmt.Printf("%s\n", me.firstName)
	fmt.Printf("%s\n", me.lastName)
	fmt.Printf("%d\n", me.age)

	beta := 20

	fmt.Printf("Default: %v\n", beta)
	fmt.Printf("Decimal: %d\n", beta)
	fmt.Printf("Binary: %b\n", beta)
	fmt.Printf("Octal: %o\n", beta)
	fmt.Printf("HexaDecimal: %x\n", beta)

	fmt.Printf("Ascii char for 56: %c\n", 56)
	fmt.Printf("Ascii char for 33: %c\n", 33)

}
