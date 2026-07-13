package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
	"strings"
)

var conferenceName = "Go Conference"

const conferenceTicket = 50

var remainingTicket uint
var slice = make(map[][string]string, 0)
var firstName string
var lastName string
var email string
var userTicket uint

func main() {
	remainingTicket = 50
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Println("We have total of", conferenceTicket, "and", remainingTicket, "are remaining")


	//if for loop is not there main will exit after booking the ticket if go routine takes longer so we add wait group
	//wg.Add(1) before goroutine to add the thread
	//wg.Wait() to wait the main function , we can add it at the end of main
	//wg.Done inside goroutine to remove the wait group means remove the thread.
	//channels is used to communicate b/w go routines
	for {

		firstNames := []string{}
		//userName = "Minakshi"
		//userTicket = 2
		fmt.Printf("Enter your first name\n")
		fmt.Scan(&firstName)

		fmt.Printf("Enter your last name\n")
		fmt.Scan(&lastName)

		fmt.Printf("ENter your email\n")
		fmt.Scan(&email)

		fmt.Printf("Enter how many tickets you want\n")
		fmt.Scan(&userTicket)

		//map
		var mapData = make(map[string]string)
		mapData["firstName"] = firstName
		mapData["lastName"] = lastName
		mapData["email"] = email
		mapData["ticket"] = strconv.FormatUint(uint64(userTicket), 10)
		validateName, validateEmail, validateTicket := helper.Validate(firstName, lastName, email, userTicket, remainingTicket)
		if validateName && validateEmail && validateTicket {
			bookTicket(firstName, lastName, userTicket, email)
			go sendTicket(firstName string, lastName string, userTicket uint, email string)
			slice = append(slice, mapData)

			for _, booking := range slice {
				//var name = strings.Fields(booking)
				firstNames = append(firstNames, booking["firstName"])

			}

			//fmt.Printf("Thank you %v for booking %v tickets you will be notified at %v\n", firstName, userTicket, email)
			fmt.Printf("Booking for the names is %v\n", firstNames)
			remainingTicket = remainingTicket - userTicket

			if remainingTicket == 0 {
				fmt.Println("No more ticket available")
				break
			}
			
		} else {
			fmt.Println("ENter new user details")
		}

	}

}


func bookTicket(firstName string, lastName string, userTicket uint, email string) {
    fmt.Printf("Thank you %v for booking %v tickets you will be notified at %v\n", firstName, userTicket, email)
}
func sendTicket(firstName string, lastName string, userTicket uint, email string) {
	time.Sleep(10* time.Second)
   var ticket = fmt.Sprintf("%v tickets are booked for %v %v user", userTicket, firstName, lastName)
   fmt.Printf("%v tickets are send to %v mail", ticket, email)
}
