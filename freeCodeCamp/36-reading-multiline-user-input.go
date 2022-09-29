package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Scanner
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Multiline Input: ")
	input := multilineReader(scanner)

	fmt.Println("Input is: ", input)
	fmt.Println("No of lines in input: ", len(input))
}

func multilineReader(scanner *bufio.Scanner) []string {
	input := []string{}

	for {
		// Scans a line from Stdin(Console)
		scanner.Scan()

		// Holds the string that scanned
		text := scanner.Text()

		if len(strings.TrimSpace(text)) != 0 {
			input = append(input, text)
		} else {
			break
		}
		fmt.Print("> ")
	}

	return input
}
