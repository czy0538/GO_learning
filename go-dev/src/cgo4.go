package main

import (
	//"C"
	"fmt"
)

/*
#include <stdio.h>
void hello()
{
	printf("hello c!");
}
*/
import "C"

type Person struct {
	name string
	age  int
	sex  byte
}

func (per Person) String() string {
	return fmt.Sprintf("%v(%v,%c)", per.name, per.age, per.sex)

}
func main() {
	p := Person{"令狐冲", 20, 'm'}
	fmt.Println(p)
	C.hello()
}
