package main

import "fmt"

func main() {
	fplus := func(x, y int) int {
		return x + y
	} //定义一个匿名函数,并将函数地址赋给fplus
	res := fplus(3, 4) //通过fplus调用匿名函数
	fmt.Println("the res is ", res)
}
