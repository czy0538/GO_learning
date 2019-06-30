package main

import "fmt"

type Birth struct {
	Year, month, day int
}
type Person struct {
	Name  string
	Age   uint
	Sex   byte //0代表男，1代表女
	Birth      //匿名组合字段
} //定义了一个名为Person的结构体类型
func main() {
	F1 := &Person{Name: "言承旭", Age: 37, Sex: 0, Birth: Birth{1977, 8, 1}}
	fmt.Println(*F1)
}
