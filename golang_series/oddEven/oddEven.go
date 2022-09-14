package main

/*
	Goal: Printing Odd numbers and Even Numbers sequentially
		1. Usage of golang channels to print odd and even numbers
		2. Understand the Logic for Odd and even Number find out
*/

import "fmt"

// Print - This helps in selecting the odd and even numbers sequentially
func Print(odd, even <-chan int) {
	for {
		select {
		case item := <-odd:
			fmt.Println(item)
		case item := <-even:
			fmt.Println(item)
		}
	}
}

func main() {
	oddChan := make(chan int)
	evenChan := make(chan int)
	go Print(oddChan, evenChan)
	for i := range [100]int{} {
		if i%2 == 0 {
			oddChan <- i
		} else {
			evenChan <- i
		}
	}
}
