package main

import "fmt"

func main() {

	pprint("string")
	pprint(23)
	pprint(23.4)

	pprint(struct {
		Name string
		Age  int
	}{
		Name: "John",
		Age:  34,
	})

}

// this is a kind of a cheat, but it works
// fmt.Printf also accept data of type any, but it prints it as a string
func pprint(x any) {
	fmt.Printf("Value is: %v\n", x)
}
