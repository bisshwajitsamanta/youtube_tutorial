package main

import "fmt"

/*
	Alternate Characters Deletion
		AAAA => returns 3 which needs to be deleted
		AABAABB => returns 2 A's which needs to be deleted

	Approach
		I need a map which says if the key found then dont add it, else add it.

*/

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
