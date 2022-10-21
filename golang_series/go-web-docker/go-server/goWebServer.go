package main

import (
	"fmt"
	"log"
	"net/http"
)

type NewCounter struct {
	CountHandler int
}

func (n *NewCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	n.CountHandler++
	ua := r.Header.Get("User-Agent")
	fmt.Fprintln(w, "Visitor Count: ", n.CountHandler)
	fmt.Fprintf(w, "User-Agent: %s\n", ua)
}

func main() {
	initialise := &NewCounter{
		CountHandler: 0,
	}
	http.Handle("/", initialise)
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
