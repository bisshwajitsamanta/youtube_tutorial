package main

import "fmt"

func main() {
	var firstNumber, secondNumber int
	var Response string
	fmt.Print("Let's add two numbers. What is your First Number ?")
	fmt.Scan(&firstNumber)
	fmt.Print("What is your second number ?")
	fmt.Scan(&secondNumber)
	result := firstNumber + secondNumber
	fmt.Printf("%d + %d=%d Did i get it right ? yes/no", firstNumber, secondNumber, result)
	fmt.Scan(&Response)
	if Response == "yes" {
		fmt.Println("I knew it, thank you")
	} else {
		fmt.Println("Oops, sorry about that.")
	}
}
