package main

import "fmt"

func main() {
	fmt.Println(m(3, 4))
}
func m(a, b int) (int, int) {
	if a > b {
		return a, b
	} else {
		return b, a
	}
}
