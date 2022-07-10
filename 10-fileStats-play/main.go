package main

import (
	"fmt"
	"os"
)


func main(){

	fmt.Println("Welcome to File Info System")
	fmt.Println("=========================")

	// fmt.Print("Enter the full path of a file:")

	// reader := bufio.NewReader(os.Stdin)
	// fileName, _ := reader.ReadString('\n')

	// fileInfo, err := os.Stat(fileName) 
	// if err != nil {
	// 	log.Fatal("Error:", err)
	// }
	// 
	fileInfo1, _ := os.Stat("myData.txt") 

	fmt.Println(fileInfo1.Size())
}