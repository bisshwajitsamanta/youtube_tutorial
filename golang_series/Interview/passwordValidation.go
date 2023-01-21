package main

/* Letters should have
1 special character
1 number
1 lower case
1 upper case
1 special character
Total Length of the password should be 6 => Done
The string contains all those and number is 6

*/

import (
	"fmt"
	"regexp"
	"strings"
)

func Validation(password string) int {
	var count int
	var (
		numeric, lowercase, uppercase, specialCharacters bool
	)

	//desired := 6
	// Check whether the length of the password is 6 or not
	//if len(password) >= desired {
	numeric = regexp.MustCompile(`\d`).MatchString(password)
	fmt.Println("Numeric:", numeric)
	lowercase = regexp.MustCompile(`[a-z]`).MatchString(password)
	fmt.Println("lowercase:", lowercase)
	uppercase = regexp.MustCompile(`[A-Z]`).MatchString(password)
	fmt.Println("uppercase:", uppercase)
	specialCharacters = strings.ContainsAny(password, "!@#$%^&*()-+")
	fmt.Println("specialCharacters:", specialCharacters)

	// Count as integer
	//if len(password) < 6 {
	//	count = 6 - len(password)
	//	return count
	//}
	if !numeric {
		count++
	}
	if !uppercase {
		count++
	}
	if !lowercase {
		count++
	}
	if !specialCharacters {
		count++
	}

	// length =47 =2<6
	// count = 3

	if len(password)+count < 6 {
		count = count + 6 - (len(password) + count)
		// 3 + 6-(2+3)
	}

	//else {
	//if len(password) < 6 {
	//charactersRequired = desired - len(password)
	//return charactersRequired
	//}
	return count

}

func main() {
	fmt.Println(Validation("47"))
}
