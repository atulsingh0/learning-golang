package helper

func ValidationCheck(userFname string, userLname string, userTickets uint, remainingTickets uint) (bool, bool) {

	isValidName := len(userFname) >= 2 && len(userLname) >= 2
	isValidTicket := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidTicket
}
