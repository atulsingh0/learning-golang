package main

import (
	"fmt"
	"time"
)

func main() {

	age := 13

	tm := time.Now()
	if age >= 18 {
		fmt.Println("You can ride.")
	}
	if age < 18 && age > 14 {
		fmt.Println("You can ride with parents.")
	}
	if age <= 14 {
		fmt.Println("You can not ride.")
	}
	fmt.Println("check independent ifs - total time:", time.Since(tm))

	// another way (with multiple independent ifs
	tm = time.Now()
	if age >= 18 {
		fmt.Println("You can ride.")
	} else if age > 14 {
		fmt.Println("You can ride with parents.")
	} else {
		fmt.Println("You can not ride.")
	}
	fmt.Println("check if else - total time:", time.Since(tm))

}
