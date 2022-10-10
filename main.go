package main

import (
	inputValidation "GoBookingApp/helper"
	"fmt"
	"strings"
	"sync"
	"time"
)


var appName  = "GoBookingApp"
const totalNumberOfTickets = 50
var remainingNumberOfTickets = 50
type bookingDetails struct{
	firstName string
	lastName string
	email string
	numberOfTicketsToBeBooked int
}

var bookings = []bookingDetails{}

var wg = sync.WaitGroup{}

func main(){
	
	greetUsers()
	
	firstName,lastName ,email,numberOfTicketsToBeBooked := input()

	isValidName,isValidEmail,isValidNumberOfTickets := inputValidation.InputValidation(firstName,lastName,email,numberOfTicketsToBeBooked,remainingNumberOfTickets)
	if(isValidName && isValidEmail && isValidNumberOfTickets){
		bookTickets(firstName,lastName,email,numberOfTicketsToBeBooked)

		wg.Add(1)
		go sendTicket(firstName,lastName,email,numberOfTicketsToBeBooked)

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
	wg.Wait()
	
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
		var names = strings.Fields(booking.firstName)
		firstNames = append(firstNames,names[0])
	
	}
	fmt.Printf("Names of people who have booked tickets: %v\n",firstNames)
}

func bookTickets(firstName string, lastName string, email string, numberOfTicketsToBeBooked int){
	remainingNumberOfTickets -= numberOfTicketsToBeBooked
	fmt.Printf("\n\n%v %v booked %v tickets.\nYou will receive a confirmation email at %v.\nThank You\n",firstName,lastName,numberOfTicketsToBeBooked,email)
	fmt.Printf("%v tickets left.\n",remainingNumberOfTickets)
	var userDetails = bookingDetails{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTicketsToBeBooked: numberOfTicketsToBeBooked,
	}
	fmt.Println("Current booking details: ",userDetails)
	bookings = append(bookings, userDetails)
	
}

func sendTicket(firstName string,lastName string,email string, numberOfTickets int){
	time.Sleep(10*time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v",numberOfTickets,firstName,lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket: \n %v \n to email: %v\n",ticket,email)
	fmt.Println("#################")
	wg.Done()
}