package main

import (
	"fmt"
)

/*
	Compare slice elements of a and b
*/

func CompareTriplets(a, b []int) []int {

	var aliceCount, bobCount int
	RecordResult := []int{}

	for i, v := range a {
		if v > b[i] {
			aliceCount += 1
		}
		if v < b[i] {
			bobCount += 1
		}
	}
	RecordResult = append(RecordResult, aliceCount, bobCount)
	return RecordResult
}

func main() {
	fmt.Println(CompareTriplets([]int{1, 2, 8}, []int{3, 4, 1}))
}
