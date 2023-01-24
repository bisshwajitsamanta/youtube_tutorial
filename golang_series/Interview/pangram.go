package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

/*
	Need to search and find all alphabets
	1. Convert the string to lower case
	2. Remove Duplicates
	3. Compare 2 slices with reflect.Deepequal or strings.containsany


*/

func Pangram(s string) string {
	//
	alphabates := []string{}
	for ch := 'a'; ch <= 'z'; ch++ {
		alphabates = append(alphabates, string(ch))
	}

	m := map[string]int{}
	addedSlice := []string{}
	removeDuplicates := []string{}
	lower := strings.ToLower(s)
	trimmed := strings.TrimSpace(lower)

	for _, v := range trimmed {
		addedSlice = append(addedSlice, string(v))
		sort.Strings(addedSlice)
	}
	for _, v := range addedSlice {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
			removeDuplicates = append(removeDuplicates, v)
		}
	}
	popped := removeDuplicates[1:]
	result := reflect.DeepEqual(alphabates, popped)
	var str string
	if result {
		str = "pangram"
	} else {
		str = "not pangram"
	}
	return str

}

func main() {
	fmt.Println(Pangram("The quick  fox jumps over the lazy dog"))
}
