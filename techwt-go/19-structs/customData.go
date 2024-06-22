package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Job    string
	gender string
}

func main() {

	atul := Person{
		Name:   "Atul",
		Age:    32,
		Job:    "Software Engineer",
		gender: "Female",
	}
	fmt.Println("Person details:", atul)

	atul.Job = "Software Architect"
	atul.gender = "Male"
	fmt.Println("Person details after change:", atul)
}
