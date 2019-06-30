package main

import "fmt"

type People struct {
	Name string
	Age  int
}

//空接口定义，可存储任意类型
type Tester interface{}

func main() {
	people := People{"张三", 20}
	it := make([]Tester, 4) //定义接口切片it，长度为4
	it[0] = 1               //整数
	it[1] = "hello"         //字符串
	it[2] = people          //结构体
	it[3] = true            //布尔
	for i, e := range it {
		if val, ok := e.(int); ok {
			fmt.Printf("it[%d] type is int,val=%d\n", i, val)
		}
		if val, ok := e.(string); ok {
			fmt.Printf("it[%d] type is string,val=%s\n", i, val)
		}
		if val, ok := e.(People); ok {
			fmt.Printf("it[%d] type is People,val=%v\n", i, val)
		}
		if val, ok := e.(bool); ok {
			fmt.Printf("it[%d] type is bool,val=%v\n", i, val)
		}
	}
}
