package main

import "fmt"

func main() {

	var (
		name           = "viktor"
		age            = 26
		salary float32 = 300_000
	)

	fmt.Printf("%s is %d years old and earning $%f per month.\n", name, age, salary)
	fmt.Printf("%s is %d years old and earning $%.0f per month.\n", name, age, salary)
	fmt.Printf("%s is %d years old and earning $%.2f per month.\n", name, age, salary)

	fmt.Printf("%T, %v\n", name, name)
	fmt.Printf("%T, %v\n", age, age)

	// let's print in another base
	fmt.Printf("%v, %d, %b, %o, %x \n", age, age, age, age, age)
	fmt.Printf("%v, %d, %b, %O, %X \n", age, age, age, age, age)

	fmt.Printf("%e\n", salary)

	fmt.Printf("%q is %d years old and earning $%.2f per month.\n", name, age, salary)

	data := 984933.124569393
	fmt.Printf("%f\n", data)

	fmt.Printf("this is my name %10s and I am learning\n", name)

	fmt.Printf("no is %012.3f", data) // 0 pad, 12 is total width, .3 is precision
}
