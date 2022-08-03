package main

import (
	"bufio"
	"fmt"
	"os"
)

type msgType int

const (
	INFO msgType = 0 + iota
	WARNING
	ERROR
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColod   = "\033[1;31m%s\033[0m"
)

func main() {

	var fileName string
	fmt.Println("welcome to file reader")

	fmt.Print("Enter the file Name: ")
	_, err := fmt.Scanf("%s", &fileName)

	if err != nil {
		showMsg(ERROR, err.Error())
	}

	file, err := os.Open(fileName)

	if err != nil {
		showMsg(ERROR, err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func showMsg(msgType msgType, msg string) {

	switch msgType {
	case INFO:
		output := fmt.Sprintf("\nINFO: %s\n", msg)
		fmt.Printf(InfoColor, output)
	case WARNING:
		output := fmt.Sprintf("\nWARNING: %s\n", msg)
		fmt.Printf(WarningColor, output)
	case ERROR:
		output := fmt.Sprintf("\nERROR: %s\n", msg)
		fmt.Printf(ErrorColod, output)
	}

}
