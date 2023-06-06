package main

import "fmt"

func main() {

	sentence := "A quick fox is jumping"
	fmt.Printf("%T\n", sentence)

	// converting it to chars (runes)
	chars := []rune(sentence)
	fmt.Println(chars) // runes are intergers (unicode)

	// to print each char/runes, convert them to string
	for _, v := range chars {
		fmt.Printf("%v -> %v, ", v, string(v))
	}
}
