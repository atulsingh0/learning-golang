package main

import "fmt"

func main() {

	ml := "this is me" +
		"who is eagar to learn go."

	fmt.Println(ml) // does not contain a new line :(

	ml = `this is me
		  who is eagar to learn go.` // contains new line and spaces

	fmt.Println(ml)
}
