package main

import (
	"fmt"
	"unicode"
)

var countItems []string

func main() {
	// How many upper letters in the word
	// Range over each letters and if rune unicode is upper it counts there.
	input := "iDontKnowYou"
	for _, v := range input {
		if unicode.IsUpper(v) {
			countItems = append(countItems, string(v))
		}
	}
	fmt.Println(len(countItems) + 1)
}
