package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
		"https://flipkart.com",
		"https://golang4.org",
		"https://golang5.org",
		"https://golang6.org",
	}

	ch1 := make(chan string)
	ch2 := make(chan string)
	for _, link := range links {
		go checkLink(link, ch1)
	}
	go func() {
		ch2 <- strings.Repeat("--", 24)
		ch2 <- "Processing List of Websites Availability:"
		ch2 <- strings.Repeat("--", 24)
	}()

	// Receive the value of Channel c
	for range links {
		select {
		case v := <-ch1:
			fmt.Println(v)
		case v := <-ch2:
			fmt.Println(v)
		case <-time.After(2 * time.Second):
			fmt.Println("timed out !")
		}
	}
}

// checkLink - This sends a message to channel about status of http Get
func checkLink(link string, c chan string) {

	if _, err := http.Get(link); err != nil {
		c <- link + " might be down"
		return
	}
	c <- link + " is up"
}
