package main

import "fmt"

func main() {
	x, y := 10, 20
	a := [...]*int{&x, &y} //指针数组，各个元素为指针
	p := &a                //指向数组地址的指针
	fmt.Printf("%T,%v\n", a, a)
	fmt.Printf("%T,%p,%v,%v,%p\n", p, p, p, *p, &p)
}
