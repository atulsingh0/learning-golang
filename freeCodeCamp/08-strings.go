package main

import "fmt"

func main() {

	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[2], s[2]) // strings are array of bytes (UTF-8)

	// strings are immutable
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)

	fmt.Printf("%v, %T\n", s+s, s+s)

}
