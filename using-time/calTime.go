package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Printf("Current time is %v\n", now)

	// date after 4 months

	getDate := now.AddDate(0, 4, 0)
	fmt.Printf("Date after 4 months: %v\n", getDate.Format("2006 Jan 02"))

	getDate = now.AddDate(0, -4, 0)
	fmt.Printf("Date 4 month ago: %v\n", getDate.Format("2006 Jan 02"))

	getHour := now.Add(2 * time.Hour)
	fmt.Printf("Time after 2 hours: %v\n", getHour.Format("15:04"))

	time.Sleep(5 * time.Second)

	fmt.Printf("Time until %v is %v\n", getHour.Format("15:04"), time.Until(getHour))

}
