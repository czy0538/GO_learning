package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		x float64 = 0.0
		y float64 = 0.0
		m int     = 0
		n int     = 1000000
		r float64 = 0.0
	)
	//fmt.Println("请输入投针次数")
	//fmt.Scanln(&n)
	rand.Seed(time.Now().Unix())
	for i := 0; i < cn; i++ {
		x = rand.Float64()
		y = rand.Float64()
		if ((x*x + y*y) - 1) <= 1e-6 {
			m++
		}
	}
	r = float64(m) / float64(n) * 4
	//r = float64(m/n) * 4 //这样结果为0
	//r = 4.0 * m / n  //报错，go的类型转换必须是显示的，不支持隐式
	fmt.Println("π = ", r)
}
