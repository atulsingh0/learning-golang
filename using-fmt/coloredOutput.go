package main

import "fmt"

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

	fmt.Println("welcome to string formatting")

	showMsg(INFO, "this is infomation")
	showMsg(ERROR, "this is Error")
	showMsg(WARNING, "this is Warning")

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
