package main

import (
	inputValidation "GoBookingApp/helper"
	"fmt"
	"strings"
)


var appName  = "GoBookingApp"
const totalNumberOfTickets = 50
var remainingNumberOfTickets = 50
var bookings = []string{}

func main(){
	
	greetUsers()
	
	for i := 0; i < 50; i++ {
		firstName,lastName ,email,numberOfTicketsToBeBooked := input()

		isValidName,isValidEmail,isValidNumberOfTickets := inputValidation.InputValidation(firstName,lastName,email,numberOfTicketsToBeBooked,remainingNumberOfTickets)
		if(isValidName && isValidEmail && isValidNumberOfTickets){
			bookTickets(firstName,lastName,email,remainingNumberOfTickets,numberOfTicketsToBeBooked)
			printNames()
			
		}else{
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

func greetUsers(){
	fmt.Printf("Welcome to %s.\n",appName)
	fmt.Println("Use this app to book tickets.")
	fmt.Printf("%v tickets available out of %v.\n",remainingNumberOfTickets,totalNumberOfTickets)
}

func input()(string,string,string,int){
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

	return firstName,lastName,email,numberOfTicketsBooked
}

func printNames(){
	firstNames := []string{}
	for _,booking:=range bookings{
		var names = strings.Fields(booking)
		firstNames = append(firstNames,names[0])
	
	}
	fmt.Printf("Names of people who have booked tickets: %v\n",firstNames)
}

func bookTickets(firstName string, lastName string, email string,remainingNumberOfTickets int, numberOfTicketsBooked int){
	remainingNumberOfTickets -= numberOfTicketsBooked
	fmt.Printf("\n\n%v %v booked %v tickets.\nYou will receive a confirmation email at %v.\nThank You\n",firstName,lastName,numberOfTicketsBooked,email)
	fmt.Printf("%v tickets left.\n",remainingNumberOfTickets)
	bookings = append(bookings, firstName+" "+lastName)
}