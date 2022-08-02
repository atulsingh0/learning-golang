package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("=================================")
	fmt.Println("Welcome to Dinner Calculation App")
	fmt.Println("=================================")

	// reading args
	args := os.Args[1:]
	// fmt.Println(len(args))

	if len(args) == 0 {
		fmt.Println("You must enter some data.")
	} else if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Usgae: ./dinnerTotal <totalAmt> <tipPercent>")
		fmt.Printf("i.e - ./dinnerTotal 100 10")
	} else {
		if len(args) != 2 {
			fmt.Println("You must enter 2 arguments! type /help for help")
		} else {
			fmt.Printf("\nCalculated Total is as below -\n")

			totalAmt, _ := strconv.ParseFloat(args[0], 32)
			tipPercent, _ := strconv.ParseFloat(args[1], 32)

			total := calculateTotal(float32(totalAmt), float32(tipPercent))
			fmt.Printf("Amount to Pay: %.2f\n", total)
			fmt.Printf("Total Tip: %.2f\n", calculateTip(float32(totalAmt), float32(tipPercent)))

		}
	}
}

func calculateTotal(amount float32, tipPercent float32) float32 {
	return amount + calculateTip(amount, tipPercent)
}

func calculateTip(amount float32, tipPercent float32) float32 {
	return amount * tipPercent / 100
}
