package main

import (
	"fmt"
	"sort"
)

type Stu struct {
	Name  string
	Num   int
	Score int
}

func main() {
	stuslice := []Stu{{"令狐冲", 201, 98}, {"岳不群", 204, 79}, {"左冷禅", 203, 88}}
	sort.Slice(stuslice, func(i, j int) bool { return stuslice[i].Score < stuslice[j].Score })
	fmt.Println(stuslice)
}
