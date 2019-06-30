// fp project main.go
package main

import (
	"fmt"
)

func main() {
	a := 3
	b := 4
	c := 0
	if a > b {
		c = a
	} else {
		c = b
	}
	fmt.Println(c)
}
