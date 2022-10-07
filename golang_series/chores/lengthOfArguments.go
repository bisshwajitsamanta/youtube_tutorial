package main

import (
	"fmt"
	"os"
)

func main() {
	code := os.Args[0]
	fmt.Println("Total Arguments Passed:", len(os.Args)-1)
	fmt.Println("Code Name:", code)
	fmt.Println("\n Arguments:")
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("\t Argument[%d]: %s\n", i, os.Args[i])
	}
}
