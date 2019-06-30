package main

import (
	"fmt"
)

/*
#include <stdio.h>
int Add(int a, int b)
 {
 int c = a+ b;
 printf("in c: result=%d\n",c);
return a + b;
 }
*/
import "C"

func main() {
	fmt.Println("")
	R := C.Add(C.int(10), C.int(20))
	fmt.Println("in go result=", R)
}
