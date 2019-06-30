package main

import "fmt"

func main() {
	var name string
	fmt.Println("input a string")
	fmt.Scanln(&name)
	fmt.Println(name)
}
