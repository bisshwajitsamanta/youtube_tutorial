package main

import "fmt"

func main() {
	var firstName, secondName string
	fmt.Print("What is your First Name ?")
	fmt.Scan(&firstName)
	fmt.Print("What is your Last Name ?")
	fmt.Scan(&secondName)

	fmt.Printf("Hello %s %s", firstName, secondName)
}
