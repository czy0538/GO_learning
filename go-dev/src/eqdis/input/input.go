package input

import "fmt"

/************************************
录入首项a1,项数num,公差dis
************************************/
func Input(a1, num, dis int) (int, int, int) {
	fmt.Println("Please input first term in a1")
	fmt.Scanln(&a1) //录入首项a
	fmt.Println("Please input term num in num")
	fmt.Scanln(&num) //录入项数n
	fmt.Println("Please input term equal difference in dis")
	fmt.Scanln(&dis) //录入公差d
	return a1, num, dis
}
