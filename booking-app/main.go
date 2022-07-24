package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

func genTicket(cityCode string, conferenceTickets uint, remainingTickets uint, bookedTickets uint) string {

	var output string
	start := conferenceTickets - remainingTickets + 1 - bookedTickets
	end := start + bookedTickets - 1

	startCode := cityCode + fmt.Sprint(start)
	endCode := cityCode + fmt.Sprint(end)

	if bookedTickets == 1 {
		output = endCode
	} else {
		output = startCode + " to " + endCode
	}
	return output
}

func calTicketCost(numOfTkts uint, ticketPrice float32) float32 {
	return float32(numOfTkts) * ticketPrice

}

func printFirstName(names []string) []string {

	var fNames []string

	for _, booking := range names { // for-each loop
		var name = strings.Fields(booking)
		fNames = append(fNames, name[0])
	}
	return fNames
}

func getUserInput() (string, string, string, uint) {
	var userFname string
	var userLname string
	var city string
	var userTickets uint

	fmt.Print("\nEnter User First Name: ")
	fmt.Scan(&userFname) // & is pointer to memory where the value is stored

	fmt.Print("Enter User Last Name: ")
	fmt.Scan(&userLname) // & is pointer to memory where the value is stored

	fmt.Print("Enter City: ")
	fmt.Scan(&city) // & is pointer to memory where the value is stored

	fmt.Print("Enter No of Ticket: ")
	fmt.Scan(&userTickets) // & is pointer to memory where the value is stored

	return userFname, userLname, city, userTickets
}

// ############## Main ###################

func main() {

	const conferenceTickets uint = 50
	var conferenceName string = "Go Conference"
	var remainingTickets uint = 50
	var ticketPrice float32 = 25.22

	var bookings = []string{"test1", "test2"}

	// fmt.Printf("conferenceName is %T, conferenceTickets is %T and remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	bookings = append(bookings, "test-user")

	// fmt.Println("Welcome to,", conferenceName, "booking portal")
	// fmt.Println("Get your ticket here to attend")

	greetUser(conferenceName, ticketPrice, conferenceTickets, remainingTickets)

	for { // infinite loop

		// var bookings [50]string // array
		// var bookings []string // unlimited size array of strings called slice

		userFname, userLname, city, userTickets := getUserInput()

		isValidName, isValidTicket := helper.ValidationCheck(userFname, userLname, userTickets, remainingTickets)
		cityCode := helper.GenCityCode(city)

		// fmt.Printf("%v, %v, %v\n", isValidName, isValidTicket, cityCode)

		name := userFname + " " + userLname

		// fmt.Printf("Memory location of userTickets : %v \n", &userTickets)

		if isValidName && isValidTicket {
			// fmt.Println(&ticketPrice) // will print the memory address

			remainingTickets = remainingTickets - userTickets
			// bookings[0] = userName // assignment to array
			bookings = append(bookings, name)

			// size := len(bookings)
			// fmt.Println(size, reflect.TypeOf(size))

			// bookings[size] = "another-test-user"

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

		// checking if tickets are 0
		if remainingTickets == 0 {
			fmt.Println("Conference is fully booked.")
			break
		}

	}
	fmt.Printf("Booking array: %v\n", bookings)
	// fmt.Printf("First Element: %v\n", bookings[0])
	fmt.Printf("First Element: %v\n", bookings[0])
	fmt.Printf("Type of Slice: %T\n", bookings)

}
