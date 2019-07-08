package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	var (
		x float64 = 0.0
		y float64 = 0.0
		m int64   = 0
	)
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000000; i++ {
		x = rand.Float64()
		y = rand.Float64()
		if  math.Pow(x, 2)+math.Pow(y, 2)-1 <= 1e-6 {
			m++
		}
	}
	fmt.Println("pi=", float64(m)/float64(1000000)*4)
}
