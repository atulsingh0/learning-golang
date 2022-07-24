package main

import "fmt"

func greetUser(conferenceName string, ticketPrice float32, conferenceTickets uint, remainingTickets uint) {
	fmt.Println("Welcome to,", conferenceName, "booking portal")
	fmt.Println("Get your ticket here to attend")
	fmt.Printf("Each Ticket Price is %v\n", ticketPrice)
	fmt.Printf("Total Tickets: %v, currently available: %v\n", conferenceTickets, remainingTickets)
}
