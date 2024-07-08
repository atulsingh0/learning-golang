package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	data := []string{
		"This is a string",
		"This is another string",
		"This is the last string",
	}

	// Printing the data
	for _, s := range data {
		fmt.Println(s)
	}

	// Saving the data to a file
	file, err := os.Create("data.out")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// One way to write to a file
	for _, s := range data {
		file.WriteString(s + "\n")
	}

	// dumping the data to a file
	combined := strings.Join(data, "\n")

	// combined data
	fmt.Printf("\nCombined data:\n%v", combined)
	err = os.WriteFile("quotes_repo.out", []byte(combined), 0666)
	if err != nil {
		fmt.Println("Error : ", err)
	}
}
