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
	F1 := &Person{}
	F1.Name = "言承旭"
	F1.Age = 37
	F1.Sex = 0
	F1.Year = 1977 //继承了Birth的Year属性，也可F1.Birth.Year = 1977
	F1.month = 8   //继承了Birth的Month属性，也可F1.Birth.month = 8
	F1.day = 1     //继承了Birth的Day属性，也可F1.Birth.day = 1
	fmt.Println(*F1)
}
