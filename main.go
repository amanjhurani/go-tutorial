package main

import "fmt"

func main()  {
	conferenceName := "Go Conference" // Syntactic Sugar for declaring variable.
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	// var bookings [50]string // array
	var bookings []string // slice

	fmt.Printf("Welcome to our %v Booking Application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are still remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Click here to collect your tickets.")

	for {
		var userName string
		var userTickets uint
		var userEmail string

		fmt.Print("Enter your first name: ")
		fmt.Scan(&userName)
		fmt.Print("Enter your email address: ")
		fmt.Scan(&userEmail)
		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if remainingTickets < userTickets {
			fmt.Println("Enter the valid remaining tickets")
			continue
		}

		remainingTickets -= userTickets
		// bookings[0] = userName // for array
		bookings = append(bookings, userName)

		fmt.Printf("Thank you %v for buying %v tickets, you will receive confirmation on email at %v.\n", userName, userTickets, userEmail)
		fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

		if remainingTickets == 0 {
			fmt.Println("All Tickets are booked, Come back next year!")
			break
		}
	}



}