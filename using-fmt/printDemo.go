package main

import "fmt"

func main() {
	var age int = 40

	fmt.Println("Jeremy is", age, "year's old")
	out, _ := fmt.Printf("Jeremy is %d year's old\n", age)

	fmt.Println("Byte written:", out)

}
