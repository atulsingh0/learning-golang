package main

import (
	"bufio"
	"fmt"
	"os"
)

func studentInfo(){
	fmt.Print("\nEnter the Student Name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n') // ignoring the returned error

	fmt.Println("Student Name: ", name)
}


func main(){

	fmt.Println("Welcome to Grading System")
	fmt.Println("=========================")

	studentInfo()
}