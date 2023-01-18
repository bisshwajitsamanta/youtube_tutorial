package main

import (
	"fmt"
	"sort"
)

type Flight struct {
	Price int
}

type byPrice []Flight

func (b byPrice) Len() int {
	return len(b)
}
func (b byPrice) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b byPrice) Less(i, j int) bool {
	return b[i].Price < b[j].Price
}

func SortbyPrice(flights []Flight) []Flight {
	sort.Sort(byPrice(flights))
	return flights
}

func main() {
	Flights := []Flight{
		{30},
		{15},
	}
	sorted := SortbyPrice(Flights)
	fmt.Println(sorted)
}
