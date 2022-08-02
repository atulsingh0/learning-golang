package main

import (
	"flag"
	"fmt"
)

func main() {

	arch := flag.String("arch", "x86", "CPU Type")

	flag.Parse()

	switch *arch {
	case "x86":
		fmt.Println("Arch type x86, Running in 32 bit mode")

	case "x64":
		fmt.Println("Arch type: x64, Running in 64 bit mode")

	case "default":
		fmt.Printf("You have enterted %v arch type which is not available\n", arch)
	}

	fmt.Println("Process Complete")
}
