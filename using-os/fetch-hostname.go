package main

import (
	"fmt"
	"os"
)

func main() {
	if host, _ := os.Hostname(); host != "" {
		fmt.Println(host)
	}
}
