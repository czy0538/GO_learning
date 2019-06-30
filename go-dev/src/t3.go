package main

import "fmt"

func makeAdder(x int) func(int) int {
	return func(y int) int { return x - y }
}
func main() {
	add5 := makeAdder(5)
	add36 := makeAdder(36)
	fmt.Println(add5(add36(1)))
}
