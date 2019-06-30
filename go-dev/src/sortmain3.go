package main

import (
	"fmt"
	"sort"
)

type StringArray []string

//实现接口Len、Less、Swap
func (p StringArray) Len() int           { return len(p) }
func (p StringArray) Less(i, j int) bool { return p[i] < p[j] }
func (p StringArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func main() {
	data := StringArray{"monday", "friday", "tuesday", "wednesday",
		"sunday", "thursday", "saturday"}
	sort.Sort(data) //排序
	for _, v := range data {
		fmt.Printf("%v\n", v) //输出
	}
}

/*sort.Sort(sort.Reverse(data)) //逆排
for _, v := range data {
	fmt.Printf("The sorted array is: %v\n", v) //输出
}*/
