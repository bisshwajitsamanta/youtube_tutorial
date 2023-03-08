package main

import "fmt"

/*
	Goal =>
		1. If the difference between Next multiple of 5 and the current number is less than 3 then round up
		2. Do not round if less than 40
		3. Do not round if the value is greater than 3 or higher
*/

// GradingStudents - Rounding Grades to next multiple of 5
func GradingStudents(grades []int) []int {
	roundOff, differenceValue := []int{}, []int{}
	var quotient int

	for _, v := range grades {
		if v%5 != 0 {
			quotient = v / 5
			quotient++
			roundOff = append(roundOff, quotient*5)
		}
	}
	// Do Slice comparison
	// Apply condition to ignore the values below 40
	for i, _ := range grades {
		if grades[i] > 40 {
			if (roundOff[i] - grades[i]) < 3 {
				fmt.Println("Element Position:", i)
				differenceValue = append(differenceValue, roundOff[i])
			}
		}
		// Remove the element from that position and append this newly element to that position
		//grades = append(grades[:i], grades[i+1:]...)
	}
	fmt.Println("Difference Slice:", differenceValue)
	fmt.Println("Original Grades:", grades)

	// [75 70 40 35]
	return nil
}

func main() {
	GradingStudents([]int{67, 73, 38, 33})
}
