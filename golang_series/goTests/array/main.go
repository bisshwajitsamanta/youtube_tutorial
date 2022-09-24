package main

func sum(numbers []int) int {
	var Sum int
	for i := range numbers {
		Sum += numbers[i]
	}
	return Sum
}
