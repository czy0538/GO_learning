package main

import "fmt"

//Go中的结构体struct地位相当于其他语言的类class
type Person struct {
	Name string
	Age  uint
	Sex  byte //0代表男，1代表女
} //定义了一个名为Person的结构体类型
/**
Go中没有构造函数的概念，
对象的创建和初始化可交由一个全局的创建函数来完成，
以NewXXX来命名，表示'构造函数'
**/
func NewPerson(name string, age uint, sex byte) Person {
	return Person{name, age, sex}
}
func main() {
	F1 := NewPerson("言承旭", 37, 0)
	fmt.Println(F1)
}
