package main

import (
	"fmt"
	"sort"
)

type day struct {
	num       int
	shortName string
	longName  string
}
type dayArray struct { //自定义星期类型
	data []day //成员为一个day切片
}

//实现接口Len、Less、Swap
func (p dayArray) Len() int           { return len(p.data) }
func (p dayArray) Less(i, j int) bool { return p.data[i].num < p.data[j].num }
func (p dayArray) Swap(i, j int)      { p.data[i], p.data[j] = p.data[j], p.data[i] }
func main() {
	Sunday := day{0, "SUN", "Sunday"}
	Monday := day{1, "MON", "Monday"}
	Tuesday := day{2, "TUE", "Tuesday"}
	Wednesday := day{3, "WED", "Wednesday"}
	Thursday := day{4, "THU", "Thursday"}
	Friday := day{5, "FRI", "Friday"}
	Saturday := day{6, "SAT", "Saturday"}
	data := []day{Tuesday, Thursday, Wednesday, Sunday, Monday, Friday, Saturday}
	a := dayArray{data}        //定义并初始化排序对象a
	sort.Sort(a)               //排序
	for _, v := range a.data { //逐个遍历输出
		fmt.Printf("%d %s %s\n", v.num, v.shortName, v.longName)
	}
}
