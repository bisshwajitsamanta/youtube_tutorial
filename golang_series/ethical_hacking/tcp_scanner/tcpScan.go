package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	ip := "scanme.nmap.org"

	//for i := 1; i < 100; i++ {
	for i := range [100]int{} {
		address := net.JoinHostPort(ip, strconv.Itoa(i))
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// Port is closed or filtered
			continue
		}
		defer conn.Close()
		fmt.Printf("[+] Connection Established ... PORT:%d %s\n", i, conn.RemoteAddr().String())
	}

}
