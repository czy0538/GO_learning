package main

import "fmt"

func main() {
	for i, v := range "Hello 世界" {
		fmt.Printf("%d,%c\n", i, v)
	}
}
