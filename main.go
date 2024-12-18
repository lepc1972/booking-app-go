package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go conference"
	const conferenceTickets int = 50 // not change
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	//fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	//fmt.Println("Welcome to our confrence booking app")
	fmt.Println("Get your tickets here to attend")

	var bookings []string //slice, abstraction of array

	//for loop infinite
	for {

		var firstName string
		var lastName string
		var email string
		var userTickets int
		println("Please insert your name")
		fmt.Scan(&firstName) //pointer
		println("Please insert your last name")
		fmt.Scan(&lastName) //pointer
		println("Please insert your email")
		fmt.Scan(&email) //pointer

		if !strings.Contains(email, "@") {
			fmt.Println("Please insert a valid email")
			fmt.Scan(&email) //pointer
		}

		println("Please insert number of tickets")
		fmt.Scan(&userTickets) //pointer

		if userTickets > int(remainingTickets) {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			println("Please insert a valid number of tickets")
			fmt.Scan(&userTickets) //pointer

		}

		//println(userName)
		//println(&userName)
		remainingTickets = remainingTickets - uint(userTickets)
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thanks  %v %v for booking %v tickets. you will receive an email confirmation at %v \n", firstName, lastName, userTickets, email)
		sendingTickets(uint(userTickets), firstName, lastName, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		//loop that iterates bookings slice, and gave us just first name
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("These are all first names of our bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			//end program
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}

	}

}

func sendingTickets(userTickets uint, firstName string, lastName string, email string) {
	//send ticket
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("################")
	fmt.Printf("Sending ticket:\n %v \n to email address %v\n", ticket, email)
	fmt.Println("################")
}
