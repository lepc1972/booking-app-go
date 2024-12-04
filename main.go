package main

import "fmt"

func main() {

	conferenceName := "Go conference"
	const conferenceTickets int = 50 // not change
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	//fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	//fmt.Println("Welcome to our confrence booking app")
	fmt.Println("Get your tickets here to attend")

}
