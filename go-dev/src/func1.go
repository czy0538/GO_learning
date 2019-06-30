package main

import "fmt"

func main() {
	f1(1, 2, 3)
	f1(1, 2, 3, 4, 5, 6)
}

//定义个变参函数，功能是输出这个变参
func f1(args ...int) { //args的实质是一个切片
	fmt.Println(args) //也可以用for range 结构逐个输出
}
