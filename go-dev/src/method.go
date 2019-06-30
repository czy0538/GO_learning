package main

import (
	"fmt"
)

type A struct {
	name string
}
type B struct {
	name string
}

func (a *A) Print() {
	a.name = "A"

}
func (b B) Print() {
	b.name = "B"

}
func main() {
	a := A{}
	b := B{}
	a.Print()
	b.Print()
	fmt.Println(a)
	fmt.Println(b)

}
