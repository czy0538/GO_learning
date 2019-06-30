package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	red := [33]int{}                //红球1~33
	blue := [16]int{}               //篮球1~16
	sred := red[0:6]                //红球选择结果
	sblue := 0                      //篮球选择结果
	for i := 0; i < len(red); i++ { //初始化
		red[i] = i + 1
		if i < len(blue) {
			blue[i] = i + 1
		}
	}
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(sred); i++ { //打散red的前六个元素
		k := i + rand.Intn(len(red)-i) //k∈[i,n-1]
		red[i], red[k] = swap(red[i], red[k])
	}
	sblue = blue[rand.Intn(16)]   //随机选择篮球
	fmt.Println(sred, ":", sblue) //输出
}
func swap(a, b int) (int, int) { //交换函数
	return b, a
}
