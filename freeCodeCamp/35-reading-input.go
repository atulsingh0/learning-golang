package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	// Creating a scanner
	scanner := bufio.NewScanner(os.Stdin)

	// Asking first no
	fmt.Print("Enter a no: ")
	scanner.Scan()
	no1 := scanner.Text()

	// Asking second no
	fmt.Print("Enter another no: ")
	scanner.Scan()
	no2 := scanner.Text()

	// Converting from string to float64
	num1, err := strconv.ParseFloat(no1, 64)
	if err != nil {
		log.Fatal("Number is not numeric: ", no1)
	}

	num2, err := strconv.ParseFloat(no2, 64)
	if err != nil {
		log.Fatal("Number is not numeric: ", no2)
	}

	// Calling hypotenuse func
	fmt.Println("Hypotenuse is:", hypotenuse(num1, num2))

}

func hypotenuse(num1, num2 float64) float64 {
	return math.Sqrt(math.Pow(num1, 2) + math.Pow(num2, 2))
}
