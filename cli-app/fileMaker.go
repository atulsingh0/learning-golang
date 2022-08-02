package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("This is a file maker application")

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Usage: ./filemaker <input_file>")
		os.Exit(3)
	} else {

		fileName := args[0]

		if _, err := os.Stat(fileName); err != nil {
			fmt.Println(err)
		} else {
			if err := os.Remove(fileName); err != nil {
				fmt.Println("File: ", fileName, "has been deleted")
			}
		}

		fmt.Println("Enter the text and press enter, type quit to exit")

		scanner := bufio.NewScanner(os.Stdin)

		file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0700)
		defer file.Close()

		if err != nil {
			fmt.Println("Err: Unable to open file: ", file.Name())
		} else {
			for scanner.Scan() {
				if scanner.Text() == "quit" {
					fmt.Println("===============================")
					contents, _ := os.ReadFile(fileName)
					fmt.Printf("%s", contents)
					fmt.Println("===============================")
					fmt.Println("Now, Quitting...")
					os.Exit(4)
				} else {
					file.WriteString(scanner.Text() + "\n")
				}
			}
		}

	}

}
