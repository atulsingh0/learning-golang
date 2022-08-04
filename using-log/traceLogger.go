package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/trace"
)

func main() {

	// Setting up tracelog
	traceFile := "tracer.log"

	file, err := os.OpenFile(traceFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0700)

	if err != nil {
		log.Fatalf("Unable to create a trace file: %v\n", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Unable to close a trace file: %v\n", err)
		}
	}()

	// Running a trace
	if err := trace.Start(file); err != nil {
		log.Fatalf("Unable to start the trace: %v\n", err)
	}
	defer trace.Stop()

	fmt.Println(rand.Intn(99) + rand.Intn(100))

	// to read the trace file, run
	// go tool trace <file>
}
