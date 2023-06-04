package main

import "fmt"

func main() {
	// strings are immutable in go

	data := "this is me"
	fmt.Printf("%s\n", data)

	// to change a letter in string, we need rune conversion

	s := []rune(data)
	s[0] = 'P'
	data = string(s)
	fmt.Println(data)

}
