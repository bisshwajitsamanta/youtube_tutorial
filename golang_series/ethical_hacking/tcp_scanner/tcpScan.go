package main

import (
	"fmt"
	"net"
	"strconv"
)

func ScanPort(port int) {
	ip := "scanme.nmap.org"
	address := net.JoinHostPort(ip, strconv.Itoa(port))
	conn, err := net.Dial("tcp", address)
	if err != nil {
		// Port is closed or filtered
		return
	}
	defer conn.Close()
	fmt.Printf("[+] Connection Established ... PORT:%d %s\n", port, conn.RemoteAddr().String())
}

func main() {

	//for i := 1; i < 100; i++ {
	for i := range [100]int{} {
		ScanPort(i)
	}

}
