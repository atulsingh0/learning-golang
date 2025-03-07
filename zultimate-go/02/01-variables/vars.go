package main

import "fmt"

func main() {

	var (
		a = 10
		b int
		c = "st-ring"
		d = true
		e = 10.10
	)

	f := float32(e) // type casting
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
	fmt.Printf("%v, %T\n", e, e)
	fmt.Printf("%v, %T\n", f, f)

}
