package main

import (
	"fmt"
	"strings"
)

func main(){

	var appName  = "GoBookingApp"
	const totalNumberOfTickets = 50
	var remainingNumberOfTickets = 50
	bookings := []string{}

	fmt.Printf("Welcome to %s.\n",appName)
	fmt.Println("Use this app to book tickets.")
	fmt.Printf("%v tickets available out of %v.\n",remainingNumberOfTickets,totalNumberOfTickets)

	for i := 0; i < 50; i++ {
		var firstName string
		var lastName string
		var email string
		var numberOfTicketsBooked int

		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets you want to book")
		fmt.Scan(&numberOfTicketsBooked)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email,"@")
		isValidNumberOfTickets := numberOfTicketsBooked > 0 && numberOfTicketsBooked<remainingNumberOfTickets

		if(isValidName && isValidEmail && isValidNumberOfTickets){
			
			remainingNumberOfTickets -= numberOfTicketsBooked
			fmt.Printf("\n\n%v %v booked %v tickets.\nYou will receive a confirmation email at %v.\nThank You\n",firstName,lastName,numberOfTicketsBooked,email)
			fmt.Printf("%v tickets left.\n",remainingNumberOfTickets)
			bookings = append(bookings, firstName+" "+lastName)
	
			firstNames := []string{}
			for _,booking:=range bookings{
				var names = strings.Fields(booking)
				firstNames = append(firstNames,names[0])
	
			}
			fmt.Printf("Names of people who have booked tickets: %v\n",firstNames)	
	
			if(remainingNumberOfTickets==0){
				fmt.Println("Sold Out!")
				break
			}	
		} else{
			if(!isValidEmail){
				fmt.Println("Invalid email")
			}
			if(!isValidName){
				fmt.Println("Invalid names")
			}
			if(!isValidNumberOfTickets){
				fmt.Println("Invalid number of tickets")
			}		
		}
		
		
		
		


		
	}
}