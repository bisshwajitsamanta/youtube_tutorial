package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	readFile, _ := ioutil.ReadFile("./doctorFile")
	data := string(readFile)
	newData := strings.Fields(strings.ReplaceAll(data, "Host:", ""))
	fmt.Println(newData)
	//fmt.Println(strings.Contains(data, "Host"))
	//if strings.Contains(data, "Host") {
	//	fmt.Println(data)
	//}
	//fmt.Println(data)

}
