package main

import "fmt"

func main() {

	// anonymous func
	func() {
		fmt.Println("Hello from anonymouns func")
	}()

	f := func() {
		fmt.Println("Ola Amigos!!")
	}
	f()

	// methods ( RUn function with context)
	// * Remember, we have calling greet method by value
	// * So if struct is big, use reference call
	gr := greeter{
		greeting: "Hi",
		name:     "Adam",
	}
	gr.greet()

}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
