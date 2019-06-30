package main

import "fmt"

func main() {
	var s, i int = 2, 1
	for i < 101 {
		s += i
		i++
	}
	fmt.Println(s)
}
