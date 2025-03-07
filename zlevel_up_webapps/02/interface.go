package main

import "fmt"

type Fruit interface {
	String() string
}

type Apple struct {
	Variety string
}

func (a Apple) String() string {
	return fmt.Sprintf("we have %s apple variety.", a.Variety)
}

type Orange struct {
	Variety string
}

func (o Orange) String() string {
	return fmt.Sprintf("we have %s orange variety.", o.Variety)
}

type Banana struct {
	Variety string
}

// function which use interface type
func printFruit(f Fruit) {
	fmt.Println(f.String())
}

func (b Banana) String() string {
	return fmt.Sprintf("we have %s banana variety.", b.Variety)
}

func main() {
	apple := Apple{Variety: "golden delicious"}
	orange := Orange{Variety: "blood"}
	banana := Banana{Variety: "plantain"}

	// we can call either individual function or interface function
	printFruit(apple)
	printFruit(orange)
	printFruit(banana)

	fmt.Println(apple.String())
}
