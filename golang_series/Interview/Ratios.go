package main

import "fmt"

func plusMinus(arr []int) {

	LengthofArray := len(arr)
	var positiveCount, negativeCount, zeroCount int

	for _, v := range arr {
		if v > 0 {
			positiveCount += 1
		}
		if v < 0 {
			negativeCount += 1
		}
		if v == 0 {
			zeroCount += 1
		}
	}
	fmt.Println("positiveCount:", float64(positiveCount)/float64(LengthofArray))
	fmt.Println("negativeCount:", float64(negativeCount)/float64(LengthofArray))
	fmt.Println("zeroCount:", float64(zeroCount)/float64(LengthofArray))
}

func main() {
	plusMinus([]int{1, 1, 0, -1, -1})
}
