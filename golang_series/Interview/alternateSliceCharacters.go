package main

import "fmt"

func RemoveCharacters(s []string) int {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] != s[j] {
				break // Not Duplicate
			} else {
				s = append(s[:i], s[j:]...)
			}
			j--
		}
	}
	return len(s)
}
func main() {
	list := []string{"a", "b", "c", "c", "c", "d"}

	fmt.Println(RemoveCharacters(list))
}
