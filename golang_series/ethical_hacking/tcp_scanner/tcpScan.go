package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

func ScanPort(port int, wg *sync.WaitGroup) {
	defer wg.Done()
	ip := "scanme.nmap.org"
	address := net.JoinHostPort(ip, strconv.Itoa(port))
	conn, err := net.Dial("tcp", address)
	if err != nil {
		// Port is closed or filtered
		// Continue you can used here if you are inside for loop else return
		return
	}
	defer conn.Close()
	fmt.Printf("[+] Connection Established ... PORT:%d %s\n", port, conn.RemoteAddr().String())
}

func main() {
	var wg sync.WaitGroup

	for i := range [65535]int{} {
		wg.Add(1)
		go ScanPort(i, &wg)
	}
	wg.Wait()
}
