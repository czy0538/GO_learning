package main

import (
	"fmt"
	"sort"
)

func main() {
	strslice := []string{"Sun", "Mon", "Tues", "Wednes", "Thur", "Fri", "Satur"}
	sort.Slice(strslice, func(i, j int) bool { return strslice[i] < strslice[j] })
	fmt.Println(strslice)
}
