package main

import "testing"

func TestSum(t *testing.T) {
	//numbers := [...]int{1, 2, 3, 4, 5}
	//got := sum(numbers)
	//want := 15
	//if got != want {
	//	t.Errorf("got %d want %d given,%v", got, want, numbers)
	//}
	t.Run("Collection of 5 Numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		got := sum(numbers)
		want := 10
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
