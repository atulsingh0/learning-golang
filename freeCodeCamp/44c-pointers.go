package main

import "fmt"

type emp struct {
	name string
	job  string
	age  int
}

func main() {

	emily := &emp{
		name: "Emily",
		job:  "CopyWriter",
		age:  29,
	}

	fmt.Printf("%+v\n", emily)
	// had birthday
	emily.birthday()

	fmt.Printf("Emily now: %+v\n", emily)

}

func (ep *emp) birthday() {
	ep.age++
}
