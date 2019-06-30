package main

import "fmt"

func main() {
	a := [...]int{1, 2}
	p := &a
	p[1] += 10 //数组指针可直接用来操作元素
	fmt.Println(p[1])
}
