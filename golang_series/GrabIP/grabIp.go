package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
	The program should grab the IP address from any API endpoint and displays on the terminal
*/

func main() {
	// Find some method to get the http request
	resp, err := http.Get("https://api.ipify.org?format=json")

	// Handle Error
	if err != nil {
		fmt.Printf("Unable to grab the data from the Api Endpoint:%w", err.Error())
	}
	data, err := ioutil.ReadAll(resp.Body)
	// Handle Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
