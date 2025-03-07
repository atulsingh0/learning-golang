package main

import "fmt"

// escape analysis
// go build -gcflags '-m -m' avoid-pointer-constructor.go

type People struct {
	Name string
	Age  int
}

func NewPeople(name string, age int) People {
	u := People{
		Name: name,
		Age:  age,
	}
	return u
}

func NewPeople2(name string, age int) *People {
	// avoid pointer constructor to create a new struct
	// this will use heap memory instead of stack memory
	u := &People{
		Name: name,
		Age:  age,
	}
	return u
}

func NewPeople3(name string, age int) *People {
	// best way to create a new struct
	return &People{
		Name: name,
		Age:  age,
	}
}

func NewPeople4(name string, age int) People {
	return People{
		Name: name,
		Age:  age,
	}
}

func main() {

	p := NewPeople("Alex", 30)
	fmt.Printf("%v, %v\n", p, &p)

	p2 := NewPeople2("Alex", 30)
	fmt.Printf("%v, %v\n", p2, &p2)

	p3 := NewPeople3("Alex", 30)
	fmt.Printf("%v, %v\n", p3, &p3)

	p4 := NewPeople4("Alex", 30)
	fmt.Printf("%v, %v\n", p4, &p4)

}
