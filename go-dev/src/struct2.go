package main

import "fmt"

type Person struct {
	Name string
	Age  uint
	Sex  byte //0代表男，1代表女
} //定义了一个名为Person的结构体类型
func main() {
	F1 := Person{}
	F1.Name = "言承旭"
	F1.Age = 37
	F1.Sex = 0
	fmt.Println(F1)
}
