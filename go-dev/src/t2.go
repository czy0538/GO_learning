package main

import "fmt"

func main() {
	a, b := 12, 16
	fmt.Println(divide(a, b))
}
func divide(x, y int) (int, int) {
	s1, s2 := x, y
	for t := x % y; t != 0; t = x % y {
		x = y
		y = t
	}
	return y, s1 * s2 / y
}
