package main

import (
	"fmt"
	"getter-setter/person"
)

func main() {

	// Creating a new person

	pn := person.Business{}

	pn.SetName("at", "ku", "sin")
	pn.SetDOB("12-12-1981")
	pn.SetContact("111-222-3333")
	pn.SetAddress("asta la vista")

	// pn.PrintInfo()

	fmt.Println("Name is: ", pn.GetName())
	fmt.Println("Contact no is: ", pn.GetContct())

}
