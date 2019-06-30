package main

import (
	"fmt"
	"math"
)

type Dis interface { //距离接口的定义
	Distance() float64 //相当于USB
}
type myfloat float64                  //相当于U盘
func (m myfloat) Distance() float64 { //实现
	if m < 0 {
		m = -m
	}
	return float64(m)
}

type Point struct { //加上相当于USB风扇
	x float64
	y float64
}

func (p Point) Distance() float64 { //实现
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func main() {
	var f myfloat = -3.14
	p := Point{3, 4}
	var d Dis
	d = f                     //插入接口，U盘连接上了主板
	fmt.Println(d.Distance()) //启动U盘，计算距离
	d = p                     //接入接口，电风扇接上主板
	fmt.Println(d.Distance()) //启动电风扇，计算距离
}
