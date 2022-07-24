package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

func bookingTicket(bookings []map[string]string, remainingTickets uint, conferenceTickets uint, ticketPrice float32) ([]map[string]string, uint) {

	var userData = map[string]string{}
	userFname, userLname, city, userTickets := getUserInput()

	isValidName, isValidTicket := helper.ValidationCheck(userFname, userLname, userTickets, remainingTickets)
	cityCode := helper.GenCityCode(city)

	if isValidName && isValidTicket {

		name := userFname + " " + userLname

		userData["firstName"] = userFname
		userData["lastName"] = userLname
		userData["city"] = cityCode
		userData["numOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

		bookings = append(bookings, userData)
		remainingTickets = remainingTickets - userTickets

		printTicket(name, cityCode, userTickets, conferenceTickets, remainingTickets, ticketPrice)
		printDetailsForAdmin(bookings, remainingTickets)

	} else if !isValidName {
		fmt.Println("Name entered are too short.")
	} else if isValidTicket {
		fmt.Printf("Can not book %v tickets more than available tickets: %v\n", userTickets, remainingTickets)
	} else {
		fmt.Println("Entered Name and Ticket count is not correct.")
	}

	return bookings, remainingTickets
}
