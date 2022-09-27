package main

import "fmt"

func main() {

	// &{var}  ==> reference pointer
	// var *{data type} ===> var is pointer type
	// *{var}  ===> dereference pointer, print value

	a := 43
	b := a
	fmt.Println(a, b)
	a = 27
	fmt.Println(a, b)

	var x int = 21
	var y *int = &x
	fmt.Println(x, y)
	x = 30
	fmt.Println(x, y) // y store the address of x
	fmt.Println(x, *y)
	*y = 14
	fmt.Println(x, *y)

	// Adding 2 to var via pointer
	fmt.Printf("Value of x: %v %v\n", x, &x)
	addTwo(&x)
	fmt.Printf("Value of x after calling addTwo func: %v %v\n", x, &x)

	// calling 2 func
	fmt.Println("=====")
	fmt.Printf("Value of x is : %v and address it is stored is %v\n", x, &x)
	// calling func getSquare
	getSquare(x)
	fmt.Printf("Value of x is after calling getSquare : %v and address it is stored is %v\n", x, &x)
	// calling func calcSquare
	calcSquare(&x)
	fmt.Printf("Value of x is after calling getSquare : %v and address it is stored is %v\n", x, &x)

	// Address of a composite data type
	fmt.Println("\nAnother example:===========================")
	k := [3]int{2, 4, 7}
	// All composite data types are by default reference to it's value,
	// hence, we are not going to see the address of var k
	fmt.Printf("Array k : %v and it's address is : %v\n", k, &k)
	j := &k[0]
	fmt.Printf("Element k[0] : %v and it's address is : %v\n", *j, j)
	fmt.Printf("Element k[1] : %v and it's address is : %v\n", k[1], &k[1])

	// Address of a composite data type
	fmt.Println("\nAnother example with composite data type:===========================")
	adam := student{
		name:     "adam",
		id:       12,
		subjects: []string{"math", "science", "computer"},
	}
	fmt.Printf("Var adam: %v and it's address: %v\n", adam, &adam)

	// initializing the address var for composite data type
	sp := &student{
		name:     "priya",
		id:       13,
		subjects: []string{"bio", "economics", "english"},
	}
	fmt.Printf("Var sp: %v and it's address: %v\n", *sp, sp)

	// initialize with new
	pr := new(student)
	fmt.Printf("Var pr: %v and it's address: %v\n", *pr, pr)
	pr.id = 14
	pr.name = "disha"
	pr.subjects = []string{"astro-physics"}
	// logically, I should use *(pr).id to point to value but Go itself will take care of that
	fmt.Printf("Var pr: %v and it's address: %v\n", *pr, pr)

	// The initial value of struct
	var st *student
	fmt.Println(st)
	// fmt.Printf("Var st: %v and it's address: %v\n", *st, st)

}

func addTwo(x *int) {
	*x += 2
}

func calcSquare(x *int) {
	fmt.Println("Address of X inside calcSquare func:", x)
	*x *= *x
}

func getSquare(x int) int {
	fmt.Println("Address of X inside getSquare func:", &x)
	return x * x
}

type student struct {
	name     string
	id       int
	subjects []string
}
