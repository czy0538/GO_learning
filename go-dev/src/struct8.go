package main

import "fmt"

type Person struct {
	Name     string
	Age      uint
	Sex      byte     //0代表男，1代表女
	Birthday struct { //成员Birthday是一个匿名结构
		Year, month, day int
	}
} //定义了一个名为Person的结构体类型
func main() {
	F1 := &Person{Name: "言承旭", Age: 37, Sex: 0} //用字面值初始化非匿名部分
	F1.Birthday.Year = 1977                     //匿名结构成员的一一赋值
	F1.Birthday.month = 8
	F1.Birthday.day = 1
	fmt.Println(*F1)
}
