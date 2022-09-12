package main

import (
	"fmt"

	sm "functions/simplemath"
)

func main() {

	//  Function Basics
	fmt.Println(sm.Add(3.4, 2.1))

	out, err := sm.Divide(223, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Restult: %0.2f\n", out)
	}

	fmt.Println(sm.Sum(2, 3, 4, 5, 6))
	fmt.Println(sm.Mul(2, 3, 4, 5, 6))
	fmt.Println(sm.Mul(2, 3, 4))

	fmt.Println("Interest on 100 by 5% for 2 year:", sm.Interest(100, 5, 1.2))

	// Calling semanticVersion function
	sv := sm.NewSemanticVersion(1, 1, 1)
	fmt.Println(sv)

	fmt.Println(sm.GetSemanticVersion(sv))
	fmt.Println(sv.String())

	sv.IncrementMajor() // this will not change the value as this is passing the value
	fmt.Println(sv.String())

	// there is an another way
	sv = sv.IncMajor()
	fmt.Println("incMajor Method:", sv)
	// But, we have to store and return the value each type

	// Reference by Pointer is a best fit for this type
	sv.IncrementMajorVer()
	fmt.Println("Ref by pointer:", sv.String())

	// Aliase Function
	func() {
		fmt.Println("\nHey there, I am from Alias world")
	}()

	// Variable as function
	greet := func() {
		fmt.Println("Hello there, How you doing??")
	}
	// calling var functon
	greet()

	// Variable function with arguement
	ola := func(name string) {
		fmt.Printf("Ola! %v, How you doing??\n", name)
	}

	ola("Atul")

	// calling from math expression
	addExp := sm.MathExp()

	fmt.Println("Math Expression:", addExp(2, 4))

	// calling stateful func

	p1 := sm.PowerOfTwo()

	value := p1()
	fmt.Println(value)
	value = p1()
	fmt.Println(value)
	value = p1()
	fmt.Println(value)

}
