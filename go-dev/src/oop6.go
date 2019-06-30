package main

import (
	"fmt"
	"math"
)

type Abser interface { //定义一个接口类型Abser
	Abs() float64 //接口中的方法名为Abs，无输入，输出类型为float64
}
type MyFloat float64 //定义了一个Myfloat类型
//Myfloat实现了接口中的Abs方法,功能为求绝对值
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct { //定义了顶点坐标类型Vertex
	X, Y float64
}

//*Vertex实现了接口的Abs方法,功能为顶点V到原点的距离
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {
	//定义变量f 初始值为-2的平方根，转换成MyFloat类型
	var f = MyFloat(-math.Sqrt2)
	//定义顶点变量V初始值为(3,4)
	v := Vertex{3, 4}
	fmt.Println(f.Abs()) //输出
	fmt.Println(v.Abs()) //输出
}
