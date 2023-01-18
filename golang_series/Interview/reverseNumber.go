package main

import "fmt"

func ReverseNumber(num int) int {
	res := 0
	for num > 0 {
		remainder := num % 10
		fmt.Println("Remainder:", remainder)
		res = (res * 10) + remainder
		fmt.Println("Res:", res)
		num /= 10
		fmt.Println("Num:", num)
		fmt.Println("======")
	}
	return res
}

func main() {
	fmt.Println(ReverseNumber(123))

}
