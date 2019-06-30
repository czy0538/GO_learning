// p4 project main.go
package main

import (
	"fmt"
)

func main() {
	var a, n, d, s int
	a, n, d = Input(a, n, d)
	s = Sum(a, n, d)
	Output(s)
}
func Input(a1, num, dis int) (int, int, int) {
	fmt.Println("input a1")
	fmt.Scanln(&a1)
	fmt.Println("input num")
	fmt.Scanln(&num)
	fmt.Println("input dis")
	fmt.Scanln(&dis)
	return a1, num, dis
}
func Sum(a1, num, dis int) int {
	return num*a1 + dis*num*(num-1)/2
}
func Output(sum int) {
	fmt.Println(sum)
}
