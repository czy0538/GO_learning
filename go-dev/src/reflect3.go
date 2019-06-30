package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Id    int
	Name  string
	Sex   bool
	Grade float32
}

func (s Student) SayHi(name string) {
	fmt.Printf("Hi %s,my name is %s.\n", name, s.Name)
}
func main() {
	stu := Student{10001, "左冷禅", true, 78.9}
	v := reflect.ValueOf(stu)
	mv := v.MethodByName("SayHi")
	//将字符串“岳不群”的Value值作为args切片的初值
	args := []reflect.Value{reflect.ValueOf("岳不群")}
	mv.Call(args) //调用
}
