package main

import "fmt"
import "runtime"

func main() {
	name := "Eric"
	go func() {
		fmt.Printf("Hello,%s", name)

	}()
	runtime.Gosched()
	name = "Harry"
}
