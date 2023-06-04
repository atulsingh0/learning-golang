package main

import (
	"fmt"
)

var (
	i int32
	j int
	k float32
	l string
)

func main() {
	i = 20
	j = 20
	k = 20
	l = "astalavista"

	fmt.Printf("%d, %T\n", i, i)
	fmt.Printf("%d, %T\n", j, j)
	fmt.Printf("%f, %T\n", k, k)
	fmt.Printf("%s, %T\n", l, l)

}
