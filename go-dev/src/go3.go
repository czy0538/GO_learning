package main

import "fmt"
import "time"

func main() {
	go fmt.Println("Go Goroutine!")
	time.Sleep(time.Millisecond)
}
