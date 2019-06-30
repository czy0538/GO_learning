package main

import "fmt"
import "time"

//接受一个参数,是只允许读取通道,除非直接强制转换
//要么只能从channel中读取数据
func sCh(ch <-chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}
func main() {
	//创建一个带100缓冲的通道 可以直接写入而不会导致主线程堵塞
	dch := make(chan int, 100)
	for i := 0; i < 100; i++ {
		dch <- i
	}
	//传递进去 只读通道
	go sCh(dch)
	time.Sleep(1e9)
}
