package main

import "fmt"
import "sort"

//自定义一个切片类型
type IntArray []int

//实现接口Len、Less、Swap
func (p IntArray) Len() int           { return len(p) }
func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }
func (p IntArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func main() {
	arr := IntArray{90, 100, 70, 80, 85, 76, 60, 88, 92, 87}
	sort.Sort(arr) //递增排序
	fmt.Println(arr)
	sort.Sort(sort.Reverse(arr)) //逆转
	fmt.Println(arr)
}
