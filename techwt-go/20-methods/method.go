package main

import "fmt"

type Student struct {
	name   string
	grades []float64
	age    int
}

// methods
func (s *Student) getGrades() []float64 {
	return s.grades
}

func (s *Student) getAverage() float64 {
	sum := 0.0
	for _, v := range s.grades {
		sum += v
	}
	return sum / float64(len(s.grades))
}

func (s *Student) changeAge(a int) {
	s.age = a
}

func main() {

	// define a student

	vicky := Student{
		name: "Vicky",
		grades: []float64{
			23.0, 11.0, 26.0, 19.0, 29.0,
		},
		age: 14,
	}

	fmt.Println("Student:", vicky)

	fmt.Println("Fetching the grades:", vicky.getGrades())
	fmt.Println("Average Marks:", vicky.getAverage())

	// changing the age
	vicky.changeAge(15)
	fmt.Println("Student:", vicky)
}
