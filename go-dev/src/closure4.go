package main

import "fmt"
import "time"

func main() {
	sign := make(chan int, 2)
	go func() {
		for i := 65; i < 75; i++ {
			fmt.Printf("%d ", i)
			time.Sleep(5 * time.Millisecond)
		}
		sign <- 1
	}()
	go func() {
		for j := 65; j < 75; j++ {
			fmt.Printf("%c ", j)
			time.Sleep(6 * time.Millisecond)
		}
		sign <- 2
	}()
	<-sign
	<-sign
}
