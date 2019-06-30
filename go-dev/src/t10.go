package main

import "fmt"

func main() {
	a := 3
	b := 4
	m := 0
	if a > b {
		m = a
	} else {
		m = b
	}
	fmt.Println(m)
}
