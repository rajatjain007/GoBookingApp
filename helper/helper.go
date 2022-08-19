package inputValidation

import (
	"strings"
)

func InputValidation(firstName string, lastName string, email string, numberOfTicketsToBeBooked int,remainingNumberOfTickets int) (bool,bool,bool){
	
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email,"@")
	isValidNumberOfTickets := numberOfTicketsToBeBooked > 0 && numberOfTicketsToBeBooked<remainingNumberOfTickets

	return isValidName,isValidEmail,isValidNumberOfTickets

}
