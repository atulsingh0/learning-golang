package main

import (
	"fmt"
	"reflect"
)

func main() {

	alpha := 4
	beta := 3.28
	gamma := "123.78"

	fmt.Println("The variable 'alpha':", alpha, "The type of var is:", reflect.TypeOf(alpha))
	fmt.Println("Converting it to float64 'alpha':", float64(alpha), "The type of var is:", reflect.TypeOf(float64(alpha)))

	fmt.Println("'beta':", beta, "The type of var is:", reflect.TypeOf(beta))
	fmt.Println("'beta' int conversion:", int(beta), "The type of var is:", reflect.TypeOf(int(beta)))

	fmt.Println("'gamma':", gamma, "The type of var is:", reflect.TypeOf(gamma))
	// fmt.Println("'beta' int conversion:", int(gamma),"The type of var is:", reflect.TypeOf(int(gamma)))
	// fmt.Println("'beta' float32 conversion:", float32(gamma),"The type of var is:", reflect.TypeOf(float32(gamma)))

}
