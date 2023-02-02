package main

import (
	"fmt"
	"reflect"
	"sort"
)

func Anagram(a, b string) bool {
	/*
		Anagram Logic says letters should be equal
		1. Range over the string and put them in a slice
		2. Compare 2 slices to check if they are anagram.
		3. Length of the two slices should be equal to be anagram.
	*/
	var anagramMeasure bool
	sliceA, sliceB := []string{}, []string{}
	for _, v := range a {
		sliceA = append(sliceA, string(v))
	}
	for _, v := range b {
		sliceB = append(sliceB, string(v))
	}

	sort.Strings(sliceA)
	sort.Strings(sliceB)

	anagramMeasure = reflect.DeepEqual(sliceA, sliceB)

	if len(sliceA) != len(sliceB) {
		return false
	}

	return anagramMeasure
}

func main() {
	fmt.Println(Anagram("evilb", "vilee"))
}
