package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	for range links {
		fmt.Println(<-c)
	}
}
func checkLink(link string, c chan string) {

	if _, err := http.Get(link); err != nil {
		c <- link + "might be down"
		return
	}
	c <- link + " is up"
}
