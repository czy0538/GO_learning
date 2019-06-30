package main
import "fmt"
import "math"
func main() {
	fmt.Println(Pi(5000))
}
func Pi(n int) float64 {
	var r float64
	ch := make(chan float64)
	for k := 0; k <= n; k++ { //并发的for循环计算每一项
		go func(kk float64) {
			ch <- math.Pow(-1, kk) / (2*kk + 1)
		}(float64(k))
	}
	for i := 0; i < n; i++ {
		r += <-ch
	}
	return 4 * r
}

/*func term(k float64, ch chan float64) {
	ch <- math.Pow(-1, k) / (2*k + 1)
}*/
/*	var t float64
	var r float64
	for k := 0; k < n; k++ {
		t = math.Pow(-1, float64(k)) / (2*float64(k) + 1)
		r += t
	}
	return 4 * r*/
