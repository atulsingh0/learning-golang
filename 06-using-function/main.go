package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(math.Floor(3.2123))

	fmt.Println(strings.Title("this is me learning golang.")) // deprecated
	fmt.Println(strings.ToTitle("this is me learning golang."))

	fmt.Println("Hi" + strings.Repeat("!", 4))

}
