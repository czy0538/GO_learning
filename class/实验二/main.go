package main

import "fmt"
import "math"

func main() {
	var n int
	fmt.Print("请输入精度:")
	_, _ = fmt.Scanln(&n)
	fmt.Println(Pi(n))
}

func Pi(n int) float64 {
	var r float64
	ch := make(chan float64, 120)
	for k := 0; k <= n; k++ {
		go func(kk float64) {
			ch <- math.Pow(-1, kk) / (2*kk + 1)
		}(float64(k))
	}
	for i := 0; i < n; i++ {
		r += <-ch
	}
	return 4 * r
}
