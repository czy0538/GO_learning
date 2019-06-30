package main

import (
	"fmt"
)

/*
#include <conio.h>
*/
import "C"

func main() {
	fmt.Println("press any key to continue")
	ch := C.getch()
	fmt.Printf("You press %c", ch)
}
