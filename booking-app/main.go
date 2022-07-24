package main

import (
	"fmt"
	"sync"
)

type userData struct {
	firstName    string
	lastName     string
	city         string
	numOfTickets uint
}

var wg = sync.WaitGroup{}

// ############## Main ###################

func main() {

	const conferenceTickets uint = 50
	var conferenceName string = "Go Conference"
	var remainingTickets uint = 50
	var ticketPrice float32 = 25.22

	// var bookings = []map[string]string{}
	var bookings = make([]userData, 0)

	greetUser(conferenceName, ticketPrice, conferenceTickets, remainingTickets)

	for { // infinite loop

		bookings, remainingTickets = bookingTicket(bookings, remainingTickets, conferenceTickets, ticketPrice)
		// checking if tickets are 0
		if remainingTickets == 0 {
			fmt.Println("\nConference is fully booked.")
			fmt.Printf("Booking array: %v\n", bookings)
			fmt.Println("Hence, exiting the program.")
			break
		}
	}
}
