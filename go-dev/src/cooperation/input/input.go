package input

import "fmt"

func Input(a int, n int, d int) (int, int, int) {
	fmt.Printf("input a、n 、d: ")
	fmt.Scan(&a, &n, &d)
	return a, n, d
}
