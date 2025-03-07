package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	val := <-ch
	fmt.Printf("\nInput Signal: %v\n", val)
	fmt.Println("exiting...")

	os.Exit(-1)
}
