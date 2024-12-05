Booking App in Go
=================

This repository contains a simple booking application implemented in Go. The application simulates a conference ticket booking system where users can book tickets and view information about the booking status.

Features
--------

*   User input for booking tickets
    
*   Displaying total tickets and available tickets
    
*   Listing booked tickets
    

Usage
-----

To run the application:

1.  Clone the repository
    
2.  Navigate to the project directory
    
3.  Run go run main.go
    

Code Structure
--------------

The main file (main.go) contains the following sections:

1.  Package declaration and imports
    
2.  Main function
    
3.  Booking struct definition
    
4.  Function to display welcome message
    
5.  Function to handle booking process
    
6.  Function to list bookings
    
7.  Function to calculate remaining tickets
    

Key Functions
-------------

### displayWelcomeMessage()

Displays a welcome message with total tickets and available tickets.

### handleBooking()

Prompts user for booking details and processes the booking.

### listBookings()

Displays all booked tickets.

### calculateRemainingTickets()

Calculates the number of remaining available tickets.

### `sendingTickets` Function in Go

The `sendingTickets` function simulates sending tickets to a user. It prints a message containing the recipient's details, the number of tickets, and the email address.

## Function Code

```go
func sendingTickets(userTickets uint, firstName string, lastName string, email string) {
	// Simulate sending tickets
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("################")
	fmt.Printf("Sending ticket:\n %v \n to email address %v\n", ticket, email)
	fmt.Println("################")
}


Contributing
------------

Contributions are welcome! Please feel free to submit a Pull Request.

License
-------

This project is licensed under the MIT License - see the [LICENSE](https://www.phind.com/LICENSE) file for details.


