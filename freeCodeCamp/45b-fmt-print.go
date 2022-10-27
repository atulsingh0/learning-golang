package main

import "fmt"

func main() {

	r := [3]rune{'a', 'b', 'c'}
	s := []string{"x", "y", "z"}

	fmt.Printf("%v\n", r)  // print the runes
	fmt.Printf("%#v\n", r) // fancy print with the details
	fmt.Printf("%T\n", r)  // print type
	fmt.Printf("%q\n", r)  // quoted chars

	fmt.Println()

	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s)
	fmt.Printf("%T\n", s)

	fmt.Println()

	m := map[string]int{"a": 23, "b": 29}

	fmt.Printf("%v\n", m)
	fmt.Printf("%#v\n", m)
	fmt.Printf("%T\n", m)

	fmt.Println()

	str := "a lazy fox jump over the brown fence"
	fmt.Printf("%v\n", str)
	fmt.Printf("%#v\n", str)
	fmt.Printf("%T\n", str)

	fmt.Println()

	bt := []byte(str)
	fmt.Printf("%v\n", bt)
	fmt.Printf("%#v\n", bt)
	fmt.Printf("%T\n", bt)

	fmt.Println()
	fmt.Printf("%v\n", string(bt))

}
