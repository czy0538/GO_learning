// p2 project main.go
package main

import (
	"fmt"
)

func main() {
	var a, b, m int = 3, 4, 0
	if a > b {
		m = a
	} else {
		m = b
	}
	fmt.Println(m)
}
