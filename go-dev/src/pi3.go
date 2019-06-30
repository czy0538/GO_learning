package main

import (
	"fmt"
	"math"
	"sync"
)

var result float64
var wg sync.WaitGroup

func f1(n float64, out chan<- float64) {
	for k := .0; k <= n; k++ {
		wg.Add(1)
		go func(k float64) {
			defer wg.Done()
			out <- math.Pow(-1, k) / (2*k + 1)
		}(k)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
}

func f2(in <-chan float64) {
	for v := range in {
		result += v
	}
	fmt.Println(result * 4)
}

func main() {
	fmt.Println("输入n值")
	var n float64

	fmt.Scanln(&n)
	ch := make(chan float64, 100)
	f1(n, ch)
	f2(ch)
}
