package main

import "fmt"

func makeadder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
func main() {
	//test := [...]string{0: "a", 1: "b", 2: "c", 3: "d", 4: "e", 6: "g"}
	//for i,v:= range test{
	//	fmt.Println(i,",",v)
	//}
	//test1:=test[:2]
	//for i,v:=range test1 {
	//	fmt.Println(i, ",", v)
	//}
	//fmt.Println(len(test))
	//fmt.Println(len(test1))
	//fmt.Println(cap(test1))

	////匿名函数1
	//f:=func (a,b int) int {
	//	return a+b
	//}
	//fmt.Println(f(2,3))
	////匿名函数2
	//sum:=func (a,b int) int {
	//	return a + b
	//}(2,3)
	//fmt.Println(sum)

	//闭包
	add := makeadder(5)
	fmt.Println(add(6))
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("有错误", err)
		} else {
			fmt.Println("正常退出")
		}
	}()
	f(101)
}
func f(i int) {
	if i > 100 {
		panic("out of bounder")
	} else {
		fmt.Println("right")
	}
}
