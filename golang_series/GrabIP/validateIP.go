package main

import (
	"fmt"
	"net"
	"reflect"
)

func main() {
	var ip string
	//inputIP := os.Args[1]
	ipv4 := "10.1.1.10005"
	fmt.Println("Enter an IP address::")
	fmt.Scanln(&ip)
	x := net.ParseIP(inputIP)
	if x != nil {
		fmt.Println("Valid Ip::", x)
		fmt.Println("Net Type::", reflect.TypeOf(x))
	} else {
		fmt.Println("InValid Ip::", ipv4)
	}
}
