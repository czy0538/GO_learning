// the telephone alphabet:
package main

import (
	"fmt"
	"sort"
)

var barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
	"delta": 87, "echo": 56, "foxtrot": 12,
	"golf": 34, "hotel": 16, "indio": 87,
	"juliet": 65, "kili": 43, "lima": 98}

func main() {
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("(%v,%v)", k, v)
	}
	keys := make([]string, len(barVal)) //定义字符串切片
	i := 0
	for k, _ := range barVal { //依次取barVal的各个关键字赋给keys[i]
		keys[i] = k
		i++
	}
	sort.Strings(keys) //对keys字符串切片排序
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys { //依次输出，
		fmt.Printf("(%v,%v)", k, barVal[k]) //barVal[k]表示取k对应的值
	}
}
