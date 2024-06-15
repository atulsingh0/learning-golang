package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	var yr string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter a year: ")
	scanner.Scan()
	yr = scanner.Text()
	fmt.Printf("Year you're are born in: %q\n", yr)

	fmt.Println("testing: ", yr)
	if no, err := strconv.ParseInt(yr, 10, 64); err == nil {
		fmt.Println("The no is:", no)
	} else {
		fmt.Printf("ParseInt: %s", err.Error())
	}
	// converting yr to int
	d, err := strconv.Atoi(yr)
	if err == nil {
		fmt.Printf("You are really born in: %d \n", d)
	} else {
		fmt.Println("Atoi: ", err.Error())
	}

	fmt.Printf("You are %d year old.\n", time.Now().Year()-d)
}
