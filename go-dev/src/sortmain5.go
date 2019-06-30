package main

import (
	"fmt"
	"sort"
)

type StringArray []string

func main() {
	data := StringArray{"monday", "friday", "tuesday", "wednesday",
		"sunday", "thursday", "saturday"}
	sort.Strings(data) //排序
	for _, v := range data {
		fmt.Printf("%v\n", v) //输出
	}
}

/*sort.Sort(sort.Reverse(data)) //逆排
for _, v := range data {
	fmt.Printf("The sorted array is: %v\n", v) //输出
}*/
