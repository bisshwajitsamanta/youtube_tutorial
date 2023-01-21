package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//type Resps struct {
//	Timestamp    string
//	Count        int
//	ResponseCode string
//}

func main() {
	file, err := os.ReadFile("./stringmanipulationset.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	data := string(file)
	m := make(map[string][]string)
	scanners := bufio.NewScanner(strings.NewReader(data))

	for scanners.Scan() {
		value := strings.SplitAfter(scanners.Text(), "responseCode")
		timeStamp, respCode := value[0][:8], value[1][:4]

		if val, ok := m[respCode]; ok {
			// Repeat value
			// Val is assigned to the value of the map[string][]string

			m[respCode] = append(val, timeStamp)

		} else {
			// First Timer
			m[respCode] = []string{timeStamp}
		}
	}
	for k, v := range m {
		fmt.Println(k, v, len(v))
	}
	//fmt.Println(m)

}
