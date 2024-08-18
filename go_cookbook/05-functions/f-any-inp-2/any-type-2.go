package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Group int
}

type Person struct {
	Name string
	Age  int
}

func main() {

	st := Student{
		Name:  "Sunny",
		Group: 12,
	}

	pe := Person{
		Name: "John",
		Age:  40,
	}

	pprint("Oh, hello there!")
	pprint(23)
	pprint(23.4)
	pprint(st)
	fmt.Println()

	//

	anypprint("Oh, hello there!")
	anypprint(23)
	anypprint(23.4)
	anypprint(st)
	anypprint([]string{"hi", "hellos"})
	fmt.Println()

	// Assertion
	assertData("hello")
	assertData(23)
	assertData(23.4)
	assertData(st)
	assetStructType(st)
	assetStructType(pe)
	anotherWay2AssetStruct(st)
	anotherWay2AssetStruct(pe)
	fmt.Println()

	// generic works best if you are looking for a specific type
	// and you know the type
	anyFuncGeneric(4)
	anyFuncGeneric(23.4)
	fmt.Println()
}

func pprint(inp any) {
	fmt.Printf("The input data is of type: %T and value: %v\n", inp, inp)
}

func anypprint(inp any) {
	fmt.Printf("Value is %v, type is %v, kind is %v\n", inp,
		reflect.TypeOf(inp), reflect.TypeOf(inp).Kind())
}

// to perform any operation on data of type "any"
// you must use type assertion to convert it to the required type
func assertData(data any) {
	switch data.(type) {
	case string:
		fmt.Printf("%v is a string\n", data)
	case int:
		fmt.Printf("%v is an int\n", data)
	case float64:
		fmt.Printf("%v is a float\n", data)
	default:
		fmt.Printf("%v is unknown data type\n", data)
	}
}

func assetStructType(data any) {
	switch data.(type) {
	case Student:
		fmt.Printf("%v is a struct %v\n", data, reflect.TypeOf(data))
	default:
		fmt.Printf("%v is unknown data type\n", data)
	}
}

func anotherWay2AssetStruct(data any) {
	std, ok := data.(Student)
	if ok {
		fmt.Printf("%v is a struct %v\n", std, reflect.TypeOf(std))
	} else {
		fmt.Printf("%v is unknown data type\n", data)
	}
}

func anyFuncGeneric[T int | float64](inp T) {
	fmt.Printf("Value is %v, type is %v, kind is %v\n", inp,
		reflect.TypeOf(inp), reflect.TypeOf(inp).Kind())
}
