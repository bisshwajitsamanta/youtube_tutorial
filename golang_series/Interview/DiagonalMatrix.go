package main

import "fmt"

/*
	1 2 3
    4 5 6
    7 8 9

	Diagonal Elements to be printed first => 1 5 9 and 4 5 7
	Sum of the elements in an array
	Subtract the value of the elements

*/

func Sum(IntSlice [][]int) int {
	var firstSection, secondSection []int
	var firstSum, secondSum, result int

	for i, v := range IntSlice {
		firstSection = append(firstSection, v[i])
	}
	for i := 0; i < len(IntSlice); i++ {
		secondSection = append(secondSection, IntSlice[i][len(IntSlice)-i-1])
	}

	for _, v := range firstSection {
		firstSum += v
	}
	for _, v := range secondSection {
		secondSum += v
	}
	if firstSum > secondSum {
		result = firstSum - secondSum
	} else {
		result = secondSum - firstSum
	}
	return result
}

func main() {

	fmt.Println(Sum([][]int{{1, 2, 2}, {4, 5, 6}, {7, 8, 9}}))

}
