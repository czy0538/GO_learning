package main

import (
	"fmt"
	"reflect"
)

func main() {
	var pi float64 = 3.14
	t := reflect.TypeOf(pi)          //获取pi的类型信息
	v := reflect.ValueOf(pi)         //获取pi的值信息
	if t.Kind() == reflect.Float64 { //t.Kind()
		fmt.Println("Type:float64", "Value:", v.Float()) //v.Float()
	}
}
