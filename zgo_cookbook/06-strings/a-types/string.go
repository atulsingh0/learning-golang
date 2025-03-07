package main

import "fmt"

func main() {

	// rune
	char := 'a'
	fmt.Println(char)

	// byte
	byteChar := byte('a')
	fmt.Println(byteChar)

	var byteChar2 byte = 'a'
	fmt.Println(byteChar2)

	// string
	str := "Hello, world!"
	fmt.Println(str)

	// multi-line string
	str2 := `
Hello, world!
This is a multi-line string.
		`
	fmt.Println(str2)
}
