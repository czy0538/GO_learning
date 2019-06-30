package main

import (
	"fmt"
	"sort"
)

var strphone = map[string]int{"Dad": 13823434432, "Mum": 13534355324, "Sun": 15263457654,
	"Sister": 18843245653, "Wife": 18823451111, "Home": 5634534}

func main() {
	keys := make([]string, len(strphone)) //定义字符串切片
	i := 0
	for k, _ := range strphone { //依次取strphone的各个关键字赋给keys[i]
		keys[i] = k
		i++
	}
	//对字符串切片keys排序
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	//依次输出
	for _, k := range keys {
		fmt.Printf("(%v,%v)", k, strphone[k]) //strphone[k]表示取k对应的值
	}
}
