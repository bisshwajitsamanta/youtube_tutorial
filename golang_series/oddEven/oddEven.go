package main

import "fmt"

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
