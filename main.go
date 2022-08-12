package main

import "fmt"

func main(){

	var appName  = "GoBookingApp"
	const totalNumberOfTickets = 50
	var remainingNumberOfTickets = 50

	fmt.Printf("Welcome to %s.\n",appName)
	fmt.Println("Use this app to book tickets.")
	fmt.Printf("%v tickets available out of %v.\n",remainingNumberOfTickets,totalNumberOfTickets)

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
	remainingNumberOfTickets -= numberOfTicketsBooked
	fmt.Printf("\n\n%v %v booked %v tickets.\nYou will receive a confirmation email at %v.\nThank You\n",firstName,lastName,numberOfTicketsBooked,email)
	fmt.Printf("%v tickets left.\n",remainingNumberOfTickets)
	


}