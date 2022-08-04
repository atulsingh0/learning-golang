package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func studentInfo() {
	fmt.Print("\nEnter the Grade: ")
	reader := bufio.NewReader(os.Stdin)
	grade, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("Fatal Error:", err)
	}

	fmt.Println("Student's grades are:", grade)
}

func main() {

	fmt.Println("Welcome to Grading System")
	fmt.Println("=========================")

	studentInfo()
}
