package main

import (
	"fmt"
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

func printFirstName(userData []map[string]string) []string {

	var fNames []string
	for _, data := range userData { // for-each loop
		fNames = append(fNames, data["firstName"])
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

	var userData = []map[string]string{}

	greetUser(conferenceName, ticketPrice, conferenceTickets, remainingTickets)

	for { // infinite loop

		userData, remainingTickets = bookingTicket(userData, remainingTickets, conferenceTickets, ticketPrice)
		// checking if tickets are 0
		if remainingTickets == 0 {
			fmt.Println("Conference is fully booked.")
			fmt.Printf("Booking array: %v\n", userData)
			break
		}
	}
}
