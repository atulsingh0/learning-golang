package main

import (
	"fmt"
	"reflect"
	"time"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
	layoutEU  = "2 January, 2006"
	layoutIST = "02/01/2006"
)

func main() {
	t := time.Now()

	fmt.Println(t)
	fmt.Println(reflect.TypeOf(t))

	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())

	yr := t.Year()
	mon := t.Month()
	day := t.Day()

	fmt.Printf("%v/%v/%v\n", yr, mon, day)

	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format(time.Kitchen))
	fmt.Println(t.Format(time.Layout))
	fmt.Println(t.Format(time.RFC1123))
	fmt.Println(t.Format(time.RFC1123Z))
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Format(time.UnixDate))

	// Custom formatting
	// Mon Jan 2 15:04:05 MST 2006
	fmt.Println(t.Format("Mon Jan 2 15:04:05 MST 2006"))
	fmt.Println(t.Format("Monday Jan 2 15:04:05 MST 2006"))
	fmt.Println(t.Format("Mon January 2 15:04:05 MST 2006"))
	fmt.Println(t.Format("15:04:05 MST, Mon, January 2, 2006"))

	fmt.Println(t.Format("Current Time is 15:04:05 MST, Today is Monday, and Date is January 2, 2006"))

	cutomDate := time.Date(1991, 12, 12, 7, 23, 56, 00, time.UTC)
	fmt.Println(cutomDate.Format("Mon Jan 2 15:04:05 MST 2006"))

	fmt.Println(cutomDate.Format(layoutEU))
	fmt.Println(cutomDate.Format(layoutISO))
	fmt.Println(cutomDate.Format(layoutUS))
	fmt.Println(cutomDate.Format(layoutIST))

}
