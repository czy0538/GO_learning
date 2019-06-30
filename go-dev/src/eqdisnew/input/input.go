package input

import (
	"fmt"
)

func Input(a1, num, dis int) (int, int, int) {
	fmt.Println("please input first term in a1")
	fmt.Scanln(&a1)
	fmt.Println("please input term num in num")
	fmt.Scanln(&num)
	fmt.Println("please input d in dis")
	fmt.Scanln(&dis)
	return a1, num, dis
}
