package main

import (
	"booking-app/helper"
	"fmt"
)

func bookingTicket(bookings []string, remainingTickets uint, conferenceTickets uint, ticketPrice float32) []string {

	userFname, userLname, city, userTickets := getUserInput()

	isValidName, isValidTicket := helper.ValidationCheck(userFname, userLname, userTickets, remainingTickets)
	cityCode := helper.GenCityCode(city)

	name := userFname + " " + userLname

	if isValidName && isValidTicket {
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, name)
		fmt.Println("---------------------------------------------------")
		fmt.Printf("User %v has booked %v tickets.\n", name, userTickets)
		fmt.Printf("Ticket Serial Nums are:  %v\n", genTicket(cityCode, conferenceTickets, remainingTickets, userTickets))
		fmt.Printf("Total amount to pay: %v\n", calTicketCost(userTickets, ticketPrice))
		fmt.Println("---------------------------------------------------")

		fmt.Printf("\nBooking Users are: %v\n", printFirstName(bookings))
		fmt.Printf("Length of Array: %v\n", len(bookings)) // problem with array is, it should have predefined size
		fmt.Printf("Total remaining tickets are : %v\n", remainingTickets)
	} else if !isValidName {
		fmt.Println("Name entered are too short.")
	} else if isValidTicket {
		fmt.Printf("Can not book %v tickets more than available tickets: %v\n", userTickets, remainingTickets)
	} else {
		fmt.Println("Entered Name and Ticket count is not correct.")
	}

	return bookings
}
