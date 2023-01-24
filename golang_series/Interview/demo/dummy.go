package main

import (
	"fmt"
	"sync"
)

// Mutexes (Mutual Exclusivity)

var msg string
var wg sync.WaitGroup

func UpdateMessage(s string) {
	defer wg.Done()

	msg = s

}

func main() {

	wg.Add(2)
	go UpdateMessage("Hello World !")
	go UpdateMessage("Hello Bisshwajit!")
	wg.Wait()
	fmt.Println(msg)
}

//func UpdateMessage(s string, m *sync.Mutex) {
//	defer wg.Done()
//	m.Lock()
//	msg = s
//	m.Unlock()
//}
//
//func main() {
//	var m sync.Mutex
//	wg.Add(2)
//	go UpdateMessage("Hello World !", &m)
//	go UpdateMessage("Hello Bisshwajit!", &m)
//	wg.Wait()
//	fmt.Println(msg)
//}
