package main

import (
	"fmt"
)

type Base struct {
	id string
}

type Person struct {
	Base
	firstname string
	lastname  string
}

type employee struct {
	Person
	salary int
}

func (recv Base) Id() string {
	return recv.id
}
func (change *Base) Setid() {
	var temp string
	fmt.Println("请输入要更改的学号")
	fmt.Scan(&temp)
	change.id = temp
}

func main() {
	employee1 := employee{Person{Base{"01"}, "zhang", "li"}, 1238}
	p := &employee1
	p.Setid()
	fmt.Println(employee1)
}
