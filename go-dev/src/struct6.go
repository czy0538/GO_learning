package main

import "fmt"

//Go中的结构体struct地位相当于其他语言的类class
type Person struct {
	Name string
	Age  uint
	Sex  byte //0代表男，1代表女
} //定义了一个名为Person的结构体类型
type Employee struct {
	Person //匿名字段
	salary int
	int         //用内置类型作为匿名字段
	Sex    byte //类似于重载
}

func main() {
	em1 := Employee{Person{"rain", 23, 0}, 5000, 100, 0}
	fmt.Println(em1)
}
