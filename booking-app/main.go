package main

import (
	"fmt"
)

// ############## Main ###################

func main() {

	const conferenceTickets uint = 50
	var conferenceName string = "Go Conference"
	var remainingTickets uint = 50
	var ticketPrice float32 = 25.22

	var bookings = []map[string]string{}

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
