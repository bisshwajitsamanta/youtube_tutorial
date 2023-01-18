package main

import (
	"fmt"
	"strings"
)

func wordcount(str string) map[string]int {
	count := map[string]int{}
	wordlist := strings.Fields(str)
	for _, v := range wordlist {
		if _, ok := count[v]; ok {
			count[v] += 1
		} else {
			count[v] = 1
		}
	}
	return count
}

func main() {
	strline := "Hello Beautiful Hello"
	for i, e := range wordcount(strline) {
		fmt.Println(i, e)
	}
}
