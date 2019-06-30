package main

import (
	"eqdis/input"
	"eqdis/output"
	"eqdis/sum"
)

func main() {
	var a, n, d, s int
	a, n, d = input.Input(a, n, d) //调用录入函数
	s = sum.Sum(a, n, d)           //调用求和函数
	output.Output(s)               //调用输出函数
}
