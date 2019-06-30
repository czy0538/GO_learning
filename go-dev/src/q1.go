package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scanln(&number)
	if number > 100 {
		number++
	} else {
		number--
	}
	fmt.Println(number)
}
