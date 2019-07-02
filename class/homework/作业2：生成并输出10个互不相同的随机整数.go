//180400501 曹志远
//2019年7月2日 21:20:28
//作业2：生成并输出10个互不相同的随机整数
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := [100]int{}
	for i := 0; i < 100; i++ {
		a[i] = i + 1
	}
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		t := i + rand.Intn(100-i)
		a[i], a[t] = a[t], a[i]
	}
	fmt.Println(a[0:10])
}
