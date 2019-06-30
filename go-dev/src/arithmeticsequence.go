package main

import "fmt"

func main() {
	var a, n, d, s int
	a, n, d = input(a, n, d) //调用录入函数
	s = sum(a, n, d)         //调用求和函数
	output(s)                //调用输出函数
}

/************************************
录入首项a1,项数num,公差dis
************************************/
func input(a1, num, dis int) (int, int, int) {
	fmt.Println("Please input first term in a1")
	fmt.Scanln(&a1) //录入首项a
	fmt.Println("Please input term num in num")
	fmt.Scanln(&num) //录入项数n
	fmt.Println("Please input term equal difference in dis")
	fmt.Scanln(&dis) //录入公差d
	return a1, num, dis
}

/************************************
等差数列求和:首项为a,公差为d,项数为n
************************************/
func sum(a1, num, dis int) int {
	return (num*a1 + num*(num-1)*dis/2)
}

/************************************
输出计算结果
************************************/
func output(r int) {
	fmt.Println("The result is ", r)
}
