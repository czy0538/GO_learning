package main

import (
	"fmt"
	"sort"
)

func main() {
	score := []int{30, 20, 50, 38, 70, 65, 23}
	//升序，闭包作为函数参数
	sort.Slice(score, func(i, j int) bool { return score[i] < score[j] })
	fmt.Println(score)
	//降序，闭包作为函数参数
	sort.Slice(score, func(i, j int) bool { return score[i] > score[j] })
	fmt.Println(score)
}
