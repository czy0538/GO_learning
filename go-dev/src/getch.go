package main

/*
#include <stdio.h>
void hello() {
    printf("hello cgo\n");
}
*/
import "C"

func main() {
	C.hello()
}
