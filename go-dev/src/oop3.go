package main

import "fmt"

type Integer int

func (a Integer) Less(b Integer) bool { //面向对象的方法
	return a < b
}
func main() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "Less 2")
	} else {
		fmt.Println(a, " Not Less 2")
	}
}
