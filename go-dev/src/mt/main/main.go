package main

import "fmt"
import "maxtwo"

func main() {
	var num1, num2, r int
	fmt.Println("input num1")
	fmt.Scanln(&num1)
	fmt.Println("input num2")
	fmt.Scanln(&num2)
	r = maxtwo.Maxtwo(num1, num2)
	fmt.Printf("the max value of num1 and num2 is %v", r)
}
