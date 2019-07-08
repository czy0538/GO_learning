package main

import "fmt"
import "time"

func main() {
	go t1()
	go t2()
	//t1,t2还没执行程序就结束了。。。
	time.Sleep(time.Second * 10)
}
func t1() {
	for i := 1; i < 6; i++ {
		fmt.Println("Hello")
		time.Sleep(time.Second)
	}
}
func t2() {
	name := [...]string{"lzy", "gcy", "hzh", "zh", "hb"}
	for _, t := range name {
		fmt.Println(t)
		time.Sleep(time.Second)
	}
}
