package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

func bookingTicket(userData []map[string]string, remainingTickets uint, conferenceTickets uint, ticketPrice float32) ([]map[string]string, uint) {

	var bookings = map[string]string{}
	userFname, userLname, city, userTickets := getUserInput()

	isValidName, isValidTicket := helper.ValidationCheck(userFname, userLname, userTickets, remainingTickets)
	cityCode := helper.GenCityCode(city)

	if isValidName && isValidTicket {

		name := userFname + " " + userLname

		bookings["firstName"] = userFname
		bookings["lastName"] = userLname
		bookings["city"] = cityCode
		bookings["numOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

		userData = append(userData, bookings)
		remainingTickets = remainingTickets - userTickets

		printTicket(name, cityCode, userTickets, conferenceTickets, remainingTickets, ticketPrice)
		printDetailsForAdmin(userData, remainingTickets)

	} else if !isValidName {
		fmt.Println("Name entered are too short.")
	} else if isValidTicket {
		fmt.Printf("Can not book %v tickets more than available tickets: %v\n", userTickets, remainingTickets)
	} else {
		fmt.Println("Entered Name and Ticket count is not correct.")
	}

	return userData, remainingTickets
}
