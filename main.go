package main

import "fmt"

func main() {
	var applicationName = "Application name"
	const confTickets = 50
	var remainingtTickets = confTickets

	fmt.Printf("Welcome to %v booking application\n", applicationName)
	fmt.Printf("We have a total of %v tickets  and %v is available\n", confTickets, remainingtTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int

	// ask user for their name

	userName = "Tom"
	userTickets = 2
	fmt.Printf("%v %v\n", userName, userTickets)

}
