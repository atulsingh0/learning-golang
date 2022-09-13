package main

import "fmt"

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal   // embedding the another struct (inheritance)
	SpeedKPH float32
	CanFly   bool
}

type Person struct {
	Name string `required min:"10"`
	Age  int
}

func main() {
	b := Bird{} // creating empty var
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48.23
	b.CanFly = false

	fmt.Println(b)

	// Another way to initialize
	c := Bird{
		SpeedKPH: 23,
		CanFly:   true,
		Animal:   Animal{Name: "Goose", Origin: "Canada"},
	}
	fmt.Println(c)

}
