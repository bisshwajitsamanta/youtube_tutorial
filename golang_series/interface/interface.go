package main

import (
	"fmt"
)

func Printout(list interface{}, v interface{}) {
	fmt.Printf(" %v :: %T\n ", list, v)
}

func main() {
	list := []interface{}{}
	list = append(list, "apple", 46, true)

	for i, v := range list {
		Printout(list[i], v)
	}
}
