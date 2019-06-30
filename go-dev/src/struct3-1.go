package main

import "fmt"

//Go中的结构体struct地位相当于其他语言的类class
type Person struct {
	Name string
	Age  uint
	Sex  byte //0代表男，1代表女
} //定义了一个名为Person的结构体类型

func main() {
	F1 := &Person{}
	F1.Name = "言承旭" //没有C中的→，依旧用.号通过指针实现间接访问
	F1.Age = 37
	F1.Sex = 0
	fmt.Println(*F1)
	fmt.Println(F1)
}
