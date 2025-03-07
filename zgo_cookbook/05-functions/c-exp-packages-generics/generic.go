package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println(sum(5, 2.4))

}

type Number interface {
	constraints.Integer | constraints.Float
}

func sum[T Number](a, b T) T {
	return a + b
}
