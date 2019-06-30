// Concurrent computation of pi.
// See http://goo.gl/ZuTZM.
//
// This demonstrates Go's ability to handle
// large numbers of concurrent processes.
// It is an unreasonable way to calculate pi.
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println(pi(5000))
}
func pi(n int) float64 {
	ch := make(chan float64)
	r := 0.0
	for k := 0; k <= n; k++ { //并行for循环
		go func(k float64) {
			ch <- 4 * math.Pow(-1, k) / (2*k + 1)
		}(float64(k))
	}

	go func() {
		for { //使用了无限循环：随着 所以term的计算完成和 ch 的变空也就结束了
			r += <-ch
		}
	}()
	time.Sleep(time.Second)
	return r
}

/*func term(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}*/

/*package main

import (
	"fmt"
	"math"
)

func main() {
	work(50000)
}
func work(n float64) float64 {
	ch := make(chan float64, 100)
	var res float64
	var i float64
	for i = 0; i <= n; i++ {
		go term(i, ch)
	}
	for i = 0; i < n; i++ {
		res += <-ch
	}
	return res
}
func term(k float64, c chan float64) {
	c <- math.Pow(-1, k) / (2*k + 1)
}
*/
