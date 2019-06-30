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

func SetValue(o interface{}) {
	v := reflect.ValueOf(o) //获取值信息
	//如果v不是指针类型或v的目标对象不可修改，则报错
	if v.Kind() != reflect.Ptr || !v.Elem().CanSet() {
		fmt.Println("cannot set!")
		return
	} else { //v是指针类型同时指向的目标对象可以修改
		v = v.Elem() //修改v为指向的目标对象
	}
	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Kind() { //获取类型,并case
		case reflect.Int:
			v.Field(i).SetInt(10001) //修改
		case reflect.String:
			v.Field(i).SetString("左冷禅")
		case reflect.Bool:
			v.Field(i).SetBool(true) //修改
		case reflect.Float32:
			v.Field(i).SetFloat(100) //修改
		}
	}
}
func main() {
	stu := Student{10002, "岳不群", false, 78.5}
	SetValue(&stu)   //传递指针，这样主调函数就可捕获修改结果
	fmt.Println(stu) //输出改变后的stu
}
