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
	test2 := test
	fmt.Println("Display test func: ", test, test2)
	test2()

	greet("Atul")

	fmt.Printf("Add: 5 + 7 = %v\n", add(5, 7))
	x, y := sqCube(5)
	fmt.Printf("Sq and Cube: 5 = %v, %v\n", x, y)
	fmt.Printf("Cube: 7 = %v\n", cube(7))

	// anonymous function

	alpha := func ()  {
		fmt.Println("calling from anynymous func: alpha")
	}

	alpha()

	beta := func (d int )  {
		fmt.Println("calling from anynymous func: beta")
		fmt.Printf("Square of %v is %v\n", d, d*d)
	}

	beta(6)

	gamma := func() string {
		return "Gamma calling"
	}()

	fmt.Println(gamma)

	delta := func(x int) string {
		return fmt.Sprintf("Value of cube of %v: %v\n", x, x*x*x)
	}(4)
	fmt.Println(delta)

	// nested function

	multi5 := getFunc(5)
	fmt.Println("Wht is multi5:", multi5)
	fmt.Println("calling multi5", multi5(7))

	multi8 := getFunc(8)
	fmt.Println("Wht is multi8:", multi8)
	fmt.Println("calling multi8", multi8(4))

	// generating a cube function
	cube := genMathFunc(3)

	fmt.Println("Cube of 2:", cube(2))
	fmt.Println("Cube of 3:", cube(3))
	fmt.Println("Cube of 4:", cube(4))

	// generating a square func

	sqr := genMathFunc(2)
	fmt.Println("Square of 2:", sqr(2))
	fmt.Println("Square of 3:", sqr(3))
	fmt.Println("Square of 4:", sqr(4))

	fmt.Println("3 ^ 4 :",genMathFunc(4)(3))

}


func getFunc( n int) func( m int) string {

	return func (m int) string  {
		return fmt.Sprintf("Value of %v * %v: %v", m, n, m * n)
	}

}

func genMathFunc (m int ) func (n  int) int {

	return func (n int) int  {
		data := 1
		for i :=0; i < m ; i++ {
			data *= n
		}
		return data
	}

}
