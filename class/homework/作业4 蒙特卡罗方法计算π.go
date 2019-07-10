package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Point struct {
	x, y float64
}

func (p Point) Length() float64 {
	return math.Pow(p.x, 2) + math.Pow(p.y, 2) - 1
}
func main() {
	var m int64 = 0
	rand.Seed(time.Now().Unix())
	var p Point
	for i := 0; i < 1000000; i++ {
		p.x = rand.Float64()
		p.y = rand.Float64()
		if p.Length() <= 1e-6 {
			m++
		}
	}
	fmt.Println("pi=", float64(m)/float64(1000000)*4)
}
