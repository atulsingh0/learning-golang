package main

import "fmt"

func main() {
	r := 'a'
	fmt.Printf("%v, %T\n", r, r)

	var p rune = 'k'
	fmt.Printf("%v, %T\n", p, p)

}
