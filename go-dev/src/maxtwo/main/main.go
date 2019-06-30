package main

import "fmt"
import "maxtwo"

func main() {
	x := 3
	y := 4
	z := 0
	z = maxtwo.Maxtwo(x, y)
	fmt.Println(z)
}
