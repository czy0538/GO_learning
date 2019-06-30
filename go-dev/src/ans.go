package main

import (
	"fmt"
)

type A struct {
	Name string
	Age  int
}
type B struct {
	Sex string
	A
}

func main() {
	var b B
	var a = A{"roy", 100}
	b = B{Sex: "man", A: a}
	fmt.Println(b, a)
}
