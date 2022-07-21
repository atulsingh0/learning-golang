package main

import (
	"fmt"
)

func main() {

	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var ticketPrice float32 = 25.22

	fmt.Println("Welcome to,", conferenceName, "booking portal")
	fmt.Println("Get your ticket here to attend")

	fmt.Printf("Total Tickets: %v, currently available: %v\n", conferenceTickets, remainingTickets)
	fmt.Printf("Each Ticket Price is %v\n", ticketPrice)

	fmt.Printf("conferenceName is %T, conferenceTickets is %T and remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	var userName string
	var userTickets uint
	// var bookings [50]string   // array
	var bookings []string

	fmt.Print("Enter User Name: ")
	fmt.Scan(&userName) // & is pointer to memory where the value is stored

	fmt.Print("Enter No of Tickets: ")
	fmt.Scan(&userTickets) // & is pointer to memory where the value is stored

	// fmt.Println(&ticketPrice) // will print the memory address

	remainingTickets = conferenceTickets - userTickets
	// bookings[0] = userName    // assignment to array
	bookings = append(bookings, userName)

	fmt.Printf("User %v has booked %v tickets.\n", userName, userTickets)
	fmt.Printf("Total remaining tickets are : %v\n", remainingTickets)

	// fmt.Printf("Booking array: %v\n", bookings)
	// fmt.Printf("First Element: %v\n", bookings[0])
	fmt.Printf("Booking Slice: %v\n", bookings)
	fmt.Printf("First Element: %v\n", bookings[0])
	fmt.Printf("Type of Slice: %T\n", bookings)
	fmt.Printf("Length of Array: %v\n", len(bookings)) // problem with array is, it should have predefined size
}
