package main

import (
	"fmt"
	"maxtwo"
)

func main() {
	var num1, num2, num3 int
	fmt.Println("Plealse input num1 by int")
	fmt.Scanln(&num1)
	fmt.Println("Plealse input num2 by int")
	fmt.Scanln(&num2)
	fmt.Println("Plealse input num3 by int")
	fmt.Scanln(&num3)
	fmt.Println("the max num is", maxtwo.Maxtwo(maxtwo.Maxtwo(num1, num2), num3))
}
