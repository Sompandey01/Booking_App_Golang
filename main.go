package main

import "fmt"

// confrenceName := "Go Confrence" // Not define const like this
    var confrenceName = "Go Confrence"
	const confrenceTickets int = 50
	var remainingTickets uint = 50
	var bookings = make([]userData, 0)

	type userData struct {
		firstName string
		lastName string
		email string
		userTickets uint
	}

func main() {
	
	greetUsers()
	
for {

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicket := validateUserInput(firstName,lastName,email,userTickets)
	
	if isValidName && isValidEmail && isValidTicket {

		bookTicket(userTickets, firstName, lastName, email)
	
    // call function print first names
	firstNames := getFirstNames()
    fmt.Printf("The first names of bookings are: %v\n", firstNames)

	// noTicketRemaining := remainingTickets == 0
	
	if remainingTickets == 0 {
		// end program
		fmt.Println("Our confrence is booked out. Come back next year.")
	 	break
	}

	} else {
		if !isValidName {
			fmt.Println("Name is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address does not contain @ sign")
		}
		if !isValidTicket{
			fmt.Println("Number of ticket you enterd is invalid")
		}
		
	}
   }
}

func greetUsers() {

	fmt.Printf("Welcome to %v booking application\n", confrenceName)
	fmt.Println("we have total of", confrenceTickets ,"tickets and", remainingTickets,"are still available")
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	 
	return firstNames
}

func getUserInput () (string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	
	// ask the user for their name 
	
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for user
	var userData = userData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)
	
	fmt.Printf("Thankyou %v %v for booking %v tickets. You'll recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
    fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, confrenceName)
	
}