package main

import "fmt"

func main() {

	x := 8
	y := &x // store x storage location
	fmt.Println("Ref of x:", y)

	// deference the variable
	fmt.Println("De-referencing the value of y:", *y)
	*y = 7
	fmt.Println("The value of x :", x, *y)

	// Calling func modify
	fmt.Println("Value of x:", x, &x)
	modify(x)
	fmt.Println("Value of x after 'modify' func:", x, &x)

	modifyRef(&x)
	fmt.Println("Value of x after 'modifyRef' func:", x, &x)

	toChange := "Hi"
	var pointer *string = &toChange

	fmt.Println("The value of 'toChange' and 'pointer': ", toChange, pointer)
	// changing
	*pointer = "Hello"
	fmt.Println("The value of 'toChange' and 'pointer' (after change): ", toChange, pointer)

}

func modify(x int) {
	fmt.Println("Pointer of x in 'modify' func:", &x)
	x = x + 2
}

func modifyRef(x *int) {
	fmt.Println("Pointer of x in 'modifyRef' func:", x)
	*x = *x + 2
}
