package main

import (
	"fmt"
	"reflect"
)

func Display(i ...interface{}) {
	fmt.Println(i)
}

func main() {
	a := "Welcome to my Youtube Channel!"
	b := 46
	c := true
	Display(a)
	Display(b)
	Display(c)
	fmt.Println(reflect.TypeOf(Display))
	Display(a, b, c)
}
