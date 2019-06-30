package main

import "fmt"

func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}
func main() {
	//定义10个通道
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int) //初始化通道i
		go Count(chs)
	}
	for _, ch := range chs {
		<-ch
	}
}
