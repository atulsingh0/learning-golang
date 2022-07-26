package main

import "fmt"

func test() {
	fmt.Println("I am running from test func.")
}

func greet(msg string) {
	fmt.Printf("Hi! %v, How are you?\n", msg)
}

func add(x int, y int) int {
	return x + y
}

func sqCube(x int) (int, int) {
	return x * x, x * x * x
}

func cube(x int) (y int) {
	y = x * x * x
	return
}

func main() {
	fmt.Println("I am running from main func.")

	test()
	greet("Atul")

	fmt.Printf("Add: 5 + 7 = %v\n", add(5, 7))
	x, y := sqCube(5)
	fmt.Printf("Sq and Cube: 5 = %v, %v\n", x, y)
	fmt.Printf("Cube: 7 = %v\n", cube(7))

}
