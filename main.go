package main

import "fmt"

func main() {
	var applicationName = "Application name"
	const confTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", applicationName)
	fmt.Printf("We have a total of %v tickets  and %v is available\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings [50]string
	var firstName string
	var lastName string
	var userTickets uint

	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter the amount of tickets you want to purchase: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole array: &v\n", bookings)
	fmt.Printf("The first value: &v\n", bookings[0])

	fmt.Printf("%v %v has %v tickets\n", firstName, lastName, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, applicationName)

}
