// hw project main.go
package main

import "fmt"

func main() {
	var a, b, c int = 3, 4, 5
	if a > b {
		c = a
	} else {
		c = b
	}
	fmt.Println(c)
}
