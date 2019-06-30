package main

import "fmt"

//Go中的结构体struct地位相当于其他语言的类class
type Person struct {
	Name string
	Age  uint
	Sex  byte //0代表男，1代表女
} //定义了一个名为Person的结构体类型
func main() {
	F1 := &Person{Name: "jow", Age: 37} //指定初始化哪些成员
	F1.Sex = 0
	fmt.Println(*F1)
}
