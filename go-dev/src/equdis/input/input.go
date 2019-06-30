package input

import "fmt"

func Input(a, n, d int) (int, int, int) {
	fmt.Println("input first term in a")
	fmt.Scanln(&a)
	fmt.Println("input n")
	fmt.Scanln(&n)
	fmt.Println("input d")
	fmt.Scanln(&d)
	return a, n, d
}
