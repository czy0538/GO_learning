package main

import (
	"fmt"
	"math"
)

//type student struct{
//	stuName string
//	age byte
//}
//type Point struct {
//	x,y float64
//}
//func (p Point) dist() float64 {
//	return math.Sqrt(p.x*p.x + p.y*p.y)
//
//}

func main() {
	//stu1:=[2]student{{"xi",16},{"he",18}}
	//fmt.Println(stu1)
	//p:=Point{3,4}
	//fmt.Println(p.dist())
	//x,y:=10,20
	//a:=[...]*int{&x,&y}
	//p:=&a
	//fmt.Printf("%T,%v\n",a,a)
	//fmt.Printf("%T,%p,%v,%v,%p\n",p,p,p,*p,&p)
	//x:=make([]int,0,5)
	//for i,v:=range x{
	//	fmt.Println(i,v)
	//}
	p := Point{3, 4}
	var tp TPoint
	tp.Point = p
	tp.z = 5

	fmt.Println(p.distance())
	fmt.Println(tp.distance())

}

type Point struct {
	x float64
	y float64
}
type TPoint struct {
	Point
	z float64
}

//func (p Point) distance() float64{
//	return math.Sqrt(p.x*p.x + p.y*p.y)
//}
//func (tp TPoint) tdis() float64  {
//	return math.Sqrt(tp.x*tp.x+tp.y*tp.y+tp.z*tp.z)
//}

type dis interface {
	distance()
}

func (p Point) distance() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}
func (tp TPoint) distance() float64 {
	return math.Sqrt(tp.x*tp.x + tp.y*tp.y + tp.z*tp.z)
}
