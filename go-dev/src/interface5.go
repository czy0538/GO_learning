package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  byte
}

func (per Person) String() string {
	return fmt.Sprintf("%v[%v,%c]", per.name, per.age, per.sex)
}
func main() {
	p := Person{"左冷禅", 120, 'm'}
	fmt.Println(p)
}
