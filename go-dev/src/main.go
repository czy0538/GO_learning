package main

import (
	"fmt"
)

func main() {
	var a, b, c int = 3, 4, 0
	if a > b {
		c = a
	} else {
		c = b
	}
	fmt.Println(c)
}
