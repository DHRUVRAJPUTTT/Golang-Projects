package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50 //uint cant let a number be negative
	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still availabe.\n", conferenceTickets, remainingTickets)
	// fmt.Println("Welcome to", conferenceName, "booking application")
	// fmt.Println("We have total of ", conferenceTickets, "tickets and", remainingTickets, "are still available. ")
	fmt.Println("Get your tickets here to attend !")

	bookings := []string{}
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		//ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of Tickets you want: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		//  isValidCity := city =="Singapore" || city =="London"

		//to fix issue of demanded tickets > available tickets
		if isValidEmail && isValidName && isValidTicketNumber {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
			// creating a Slice(dynamic array) of name firstNames to print only 1st name of bookings

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				if len(names) > 0 {
					firstNames = append(firstNames, names[0])
				}

			}
			fmt.Printf("The first names of bookings are %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Our conference is booked out. Come back next year. ")
				break

			}
			// if  remainingTickets ==0 {
			// 	//end the program
			// 	fmt.Println("Our conference is booked out. Come back next year. ")
			// 	break

			// }
		} else {
			if !isValidName {
				fmt.Println("First name or Last name you entered is too short ")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid ")
			}

		}

	}

}
