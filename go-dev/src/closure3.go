package main

import "fmt"

func f() (ret int) {
	defer func() {
		ret++
	}() //定义并在return 1后调用匿名函数
	return 1
}
func main() {
	fmt.Println(f())
}
