package main

import "fmt"
import "time"

func main() {
	var sum int = 0
	t0 := time.Now() //获取当前时间
	for i := 0; i < 100000; i++ {
		sum += i
	}
	t := time.Now().Sub(t0) //获取当前时间和t0做差
	fmt.Println(t)
}
