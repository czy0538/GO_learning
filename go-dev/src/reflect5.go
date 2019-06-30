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
type Monitor struct {
	Student
	As string
}

func main() {
	stu := Monitor{Student{10002, "岳不群", true, 78.9}, "班长"}
	t := reflect.TypeOf(stu)
	v := reflect.ValueOf(stu)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Anonymous { //如果是匿名字段
			for j := 0; j < v.Field(i).NumField(); j++ {
				//通过value的FieldByIndex()方法获取嵌入字段的索引路径 []int{i,j},然后取值
				fmt.Println("Embedded fields:", v.FieldByIndex([]int{i, j}).Interface())
			}
		} else {
			fmt.Println("Normal fields:", v.Field(i).Interface())
		}
	}
}
