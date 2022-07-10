package main

import "fmt"

func main() {
	var name string
	var mName, lName string
	var age int
	var salary int64
	var height float32
	var siblings, sAge int = 1, 2
	var address = "200 Queen St., Toronto"

	var school string // No value assigned
	var grade int     // No value assigned

	name = "Alexa"
	mName, lName = "maria", "D'souza"
	age = 8
	salary = 10101010
	height = 3.25

	fatherName := "Peter Parker" // Short var declarations

	fmt.Println("Hi, My name is", name, "and I am", age, "years old.")
	fmt.Println("My middle name is", mName, "and last is", lName)
	fmt.Println("I have", siblings, "brother and his age is", sAge)
	fmt.Println("My address is", address)
	fmt.Println("My school is", school, "and I am in grade", grade) // school = "" and grade = 0

	fmt.Println("My Dad's salary is", salary, "and his height is", height, "feet.")
	fmt.Println("My Dad's name is", fatherName)

}
