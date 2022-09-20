package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("start")
	fmt.Println("middle")
	fmt.Println("end")

	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")

	// defer keyword follow LIFO

	fh, err := os.Open("/etc/hosts")

	// defer statement execute just before main exit
	defer fh.Close()

	if err != nil {
		fmt.Println("File read err:", err)
	} else {
		fmt.Println(fh.Name())
	}

	// defer execute at last but initialized wherever it is defined.
	a := 5
	defer fmt.Println(a)
	a = 7
}
