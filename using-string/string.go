package main

import "fmt"

func main() {

	str := "Go is awesome!"

	fmt.Println(str)
	// fmt.Printf("%x\n", str)

	for i := 0; i < len(str); i++ {
		fmt.Printf("%x %c\n", str[i], str[i])
	}

	fmt.Println(str[3])  // return byle value of 4th index
	fmt.Println(str[:4]) // slice will return the actual value
}
