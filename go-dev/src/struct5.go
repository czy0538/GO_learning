package main

import "fmt"

//Go中的结构体struct地位相当于其他语言的类class
type Person struct {
	Name string
	Age  uint
	Sex  byte //0代表男，1代表女
} //定义了一个名为Person的结构体类型
func main() {
	F1 := new(Person) //新建一个Person指针对象，内容空
	F1.Name = "言承旭"
	F1.Age = 37
	F1.Sex = 0
	fmt.Println(*F1)
}
