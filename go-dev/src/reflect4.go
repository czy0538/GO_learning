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

func (s Student) SayHi() {
	fmt.Printf("Hi,nice to meet you!\n")
}
func (s Student) Myname() {
	fmt.Printf("My name is %s,I come from China.", s.Name)

}
func StructInfo(o interface{}) {
	t := reflect.TypeOf(o)
	//判断是否为结构体
	if k := t.Kind(); k != reflect.Struct {
		fmt.Printf("This value is not a struct,It's %v.", k)
		return
	}
	fmt.Println("Struct name is", t.Name())
	fmt.Println("Fields of the struct is:")
	v := reflect.ValueOf(o)
	//获取字段Type和Value值
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		fmt.Printf("%6s:%v=%v\n", field.Type, field.Name, value)
	}
	fmt.Println("Methods of the struct is:")
	//获取方法name和Type信息
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("%6s:%v\n", method.Name, method.Type)
	}
}
func main() {
	stu := Student{10002, "岳不群", true, 78.9}
	StructInfo(stu)
}
