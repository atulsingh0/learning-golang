package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	start := time.Now()
	bufioScanner("/var/log/install.log")
	fmt.Println("Total time taken via bufio: ", time.Since(start))

}

func bufioScanner(file string) {

	// Reading via bufio Scanner
	// Read huge file line by line
	fh, err := os.Open(file)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer fh.Close()

	scanner := bufio.NewScanner(fh)

	for scanner.Scan() {
		scanner.Text()
	}
}
