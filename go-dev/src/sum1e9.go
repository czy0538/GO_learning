package main

import "fmt"

func main() {
	var sum int64
	for i := 0; i < 100000000; i++ {
		sum += int64(i + 1)
	}
	fmt.Println(sum)
}
