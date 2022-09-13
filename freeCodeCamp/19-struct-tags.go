package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `required max:10` // TODO: Take a look in the syntax
	Origin string
}

func main() {

	anm := Animal{
		Name:   "this is calling care",
		Origin: "Canada",
	}
	fmt.Println(anm)
	fmt.Println(anm.Origin)

	t := reflect.TypeOf(anm)
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
