package main

import (
	"fmt"
)

func maxfactor(a, b int) int {
	if a > b {
		a, b = b, a
	}
	var c int
	for i := 2; i < a; i++ {
		if b%i == 0 {
			c = i
		}
	}
	return c
}

func main() {
	var (
		a, b, c int
	)
	//fmt.Scanln(&a, &b)
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	c = maxfactor(a, b)
	fmt.Println(a * b / c)
}
