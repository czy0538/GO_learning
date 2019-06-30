package main

import "fmt"
import "sort"

func main() {
	arr := []int{90, 100, 70, 80, 85, 76, 60, 88, 92, 87}
	sort.Ints(arr) //升序
	fmt.Println(arr)
	sort.Sort(sort.Reverse(sort.IntSlice(arr))) //降序
	fmt.Println(arr)
}
