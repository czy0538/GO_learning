package main

import "fmt"

func main() {
	F1 := struct { //定义了一个匿名结构，属性也是匿名的
		string //姓名
		uint   //年龄
		byte   //0代表男，1代表女
	}{"言承旭", 37, 0} //按照顺序初始化各个匿名字段
	fmt.Println(F1)
}
