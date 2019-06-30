package main

import (
	"fmt"
	"math"
)

func main() {
	done := make(chan bool)
	data := make(chan float64)

	go cal(data)
	go res(data, done)

	<-done
}

func cal(data chan float64) {
	for i := 0; i < 10000; i++ {
		k := float64(i)
		data <- math.Pow(-1, k) / (2*k + 1)
	}
	close(data)
}

func res(data chan float64, done chan bool) {
	var result float64
	for x := range data {
		result += x
	}
	fmt.Println("the result:", 4*result)
	done <- true
}
