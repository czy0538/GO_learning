package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	a = 3
	b = 4
	c = 0
	if a > b {
		c = a
	} else {
		c = b
	}
	fmt.Println(c)

}
