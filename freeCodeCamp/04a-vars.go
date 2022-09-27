package main

import "fmt"

func main() {

	alpha := 0
	var beta int
	var gamma uint64 = 0

	var zeta string
	delta := "hi"

	fmt.Println(alpha)
	fmt.Println(beta)
	fmt.Println(gamma)
	fmt.Println(zeta)
	fmt.Println(delta)

	var (
		i int
		j float64
		k string = "this"
	)

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
}
