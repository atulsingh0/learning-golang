package main

import "fmt"

func main() {
	fmt.Println(square(4))
	fmt.Println(square(4.2))

	fmt.Println(square2(4.0))
	fmt.Println(square2(4.2))

}

// use of generic type
// square function can take int or float64 as argument and return int or float64
// square[T int | float64] means square function can take any type T which is either int or float64
// golang will automatically convert int to float64 when calling square function
func square[T int | float64](x T) T {
	return x * x
}

// another way to write square function
type Number interface {
	~int | ~float64
}

func square2[T Number](x T) T {
	return x * x
}
