package main

import "fmt"

func main() {

	grades1 := 92
	grades2 := 91
	grades3 := 90
	grades4 := 89

	fmt.Printf("Grades are: %v, %v, %v, %v\n", grades1, grades2, grades3, grades4)

	// Array can store only 1 type of data
	// Length of Array must be pre-defined
	// Either by [No] way or [...]

	grades := [4]int{92, 91, 90, 89}
	fmt.Printf("Grades are: %v\n", grades)

	grade := [...]int{92, 91, 90, 89}
	fmt.Printf("Grades are: %v\n", grade)

	var students [3]string // empty string
	fmt.Printf("Students: %v \n", students)

	students[0] = "Akash"
	students[1] = "Vikash"
	fmt.Printf("Students: %v \n", students)
	fmt.Printf("Length of Students: %v \n", len(students))
	fmt.Printf("Capacity of Students: %v \n", cap(students))

	a := [3]int{1, 2, 4}
	b := a  // Copying an Array creates another copy array
	c := &a // This is referencing the same data as a

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n\n", c)

	a[1] = 99
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
}
