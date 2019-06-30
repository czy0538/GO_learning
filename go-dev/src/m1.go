package main

import "fmt"

func Input(a1, num, dis int) (int, int, int) {
	fmt.Println("please input first term in a1")
	fmt.Scanln(&a1)
	fmt.Println("please input term num in num")
	fmt.Scanln(&num)
	fmt.Println("please input d in dis")
	fmt.Scanln(&dis)
	return a1, num, dis

}
func Sum(a, n, d int) int {
	return n*a + n*(n-1)*d/2
}
func Output(r int) {
	fmt.Println("The result is", r)

}
func main() {
	var a, n, d, s int
	a, n, d = Input(a, n, d)
	s = Sum(a, n, d)
	Output(s)
}
