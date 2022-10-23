package main

import "fmt"

func main() {

	// pointers are the address of a memory location

	no := 5
	ads := &no // add is a variable and holds the address of no

	no_2 := *ads // no_2 is a variable and holds the "value" which is placed at the address holds by ads variable

	fmt.Println("no:", no)
	fmt.Println("Address of no:", &no, ads)
	fmt.Printf("Value at memory address %v: %v\n\n", ads, no_2)

	if no == no_2 {
		fmt.Println("Numbers are equal")
	} else {
		fmt.Println("Not equal.")
	}

	fmt.Println("Value of *&no:", *&no)
	fmt.Printf("Type of an address variable: %T\n", ads) // * prefix a type denotes a pointer type

	fName := "Emily"
	ref := &fName

	fmt.Printf("\nThe type of var 'ref' is: %T & memory address of var 'ref' is %v\n", ref, &ref)
	fmt.Println("var 'ref' holds: ", ref)
	fmt.Println("The value at memory hold by var 'ref':", *ref)

	lName := "Dsouza"
	ref = &lName
	fmt.Println("Now, the value at memory hold by var 'ref':", *ref)
	fmt.Println("var 'ref' holds: ", ref)
	fmt.Println("Memory address of var 'ref' is still:", &ref)

	// ref variable is previously pointing to var fName address and now to var lName address

	*ref = "Ramsay" // changing the value at memory address hold by ref var

	fmt.Println("The value is now: ", *ref, lName) // both got changed

}
