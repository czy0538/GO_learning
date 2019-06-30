package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	a := Person{name: "joe", age: 12}
	fmt.Println(a)
	A(a)
	fmt.Println(a)

}
func A(p Person) {
	p.age = 13
	fmt.Println(p)

}
