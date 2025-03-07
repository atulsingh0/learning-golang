package main

import "fmt"

func main() {
	defer fmt.Println("defer on main")
	fmt.Println("main")
	A()
	fmt.Println("end of main")
}

func A() {
	defer fmt.Println("defer on A")
	fmt.Println("A")
	B()
	fmt.Println("end of A")
}

func B() {
	defer dontPanic()
	fmt.Println("B")
	C()
	fmt.Println("end of B")
}

func C() {
	defer fmt.Println("defer on C")
	fmt.Println("C")
	D()
	fmt.Println("end of C")
}

func D() {
	defer fmt.Println("defer on D")
	fmt.Println("D")

	// when the panics is caught, it will be executed from the latest defer
	// and break the normal execution flow
	panic("panic on D")
	fmt.Println("end of D")
}

func dontPanic() {
	err := recover()
	if err != nil {
		fmt.Println("panic is called: ", err)
	} else {
		fmt.Println("defer on B")
	}
}
