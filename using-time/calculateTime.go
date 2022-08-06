package main

import (
	"fmt"
	"time"
)

func main() {
	first := time.Now()
	// fmt.Printf("Current time is %v\n", first)
	fmt.Printf("Current time is %v\n", first.Format("15:04:05"))

	// time.Sleep(2 * 1000000000) // Nano-seconds

	time.Sleep(2 * time.Second)

	second := time.Now()
	fmt.Printf("Now time is %v\n", second.Format("15:04:05"))

	cusDate := time.Date(2010, 07, 8, 9, 00, 00, 00, time.UTC)

	elapsed := time.Since(cusDate)
	fmt.Printf("Total time passed since %v is %v\n", cusDate.Format("2006-01-02"), elapsed)

	fmt.Printf("Total time passed since %v is %v hours, %v minutes and %v seconds\n", cusDate.Format("2006-01-02"), elapsed.Hours(), elapsed.Minutes(), elapsed.Seconds())
	fmt.Printf("Total time passed since %v is %.2f hours, %.2f minutes and %.2f seconds\n", cusDate.Format("2006-01-02"), elapsed.Hours(), elapsed.Minutes(), elapsed.Seconds())
	fmt.Printf("Total time passed since %v is %.0f hours, %.0f minutes and %.0f seconds\n", cusDate.Format("2006-01-02"), elapsed.Hours(), elapsed.Minutes(), elapsed.Seconds())
}
