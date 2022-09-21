package main

import (
	"fmt"
	"runtime"
)

// 1e9 is 100 million
var listLength int = 1e9

type memoryStats struct {
	allocated      string
	totalAllocated string
	system         string
}

// getMemory finds the allocated, totalAllocated and system memory
func getMemory() memoryStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return memoryStats{
		allocated:      fmt.Sprintf("%d MiB", m.Alloc/1024/1024),
		totalAllocated: fmt.Sprintf("%d MiB", m.TotalAlloc/1024/1024),
		system:         fmt.Sprintf("%d MiB", m.Sys/1024/1024),
	}
}

func main() {
	list := []int{}
	fmt.Println("List length:", len(list))
	fmt.Printf("Memory usage: %+v\n", getMemory())

	for i := 0; i < listLength; i++ {
		list = append(list, i)
		//fmt.Println(list)
	}

	fmt.Println("List length:", len(list))
	fmt.Printf("Memory usage: %+v\n", getMemory())
}
