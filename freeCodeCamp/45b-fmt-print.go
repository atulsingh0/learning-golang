package main

import "fmt"

func main() {

	r := [3]rune{'a', 'b', 'c'}
	s := []string{"x", "y", "z"}

	fmt.Printf("%v\n", r)  // print the runes
	fmt.Printf("%#v\n", r) // fancy print with the details (go styles)
	fmt.Printf("%T\n", r)  // print type
	fmt.Printf("%q\n", r)  // quoted chars

	fmt.Println()

	fmt.Printf("%v\n", s)  // string array
	fmt.Printf("%#v\n", s) // go style string array
	fmt.Printf("%T\n", s)  // data type

	fmt.Println()

	m := map[string]int{"a": 23, "b": 29}

	fmt.Printf("%v\n", m)  // map
	fmt.Printf("%#v\n", m) // go style map
	fmt.Printf("%T\n", m)  // data type

	fmt.Println()

	str := "a lazy fox jump over the brown fence"
	fmt.Printf("%v\n", str)
	fmt.Printf("%#v\n", str)
	fmt.Printf("%T\n", str)

	fmt.Println()

	bt := []byte(str)       // convert string to bytes
	fmt.Printf("%v\n", bt)  // int8 (byte) for each character
	fmt.Printf("%#v\n", bt) // go style byte
	fmt.Printf("%T\n", bt)  // bytes are int8

	fmt.Println()
	fmt.Printf("%v\n", string(bt)) // convert byte array to strings

}
