package main

import (
	"fmt"
)

type Stu struct {
	stunum   string
	stuname  string
	stuscore int
}

func main() {
	var key string
	stumap := make(map[string]Stu)
	stumap["160400101"] = Stu{"160400101", "lhc", 95}
	stumap["160400102"] = Stu{"160400102", "mcf", 100}
	fmt.Println("请输入查找的学号")
	fmt.Scanln(&key)
	stuinf, ok := stumap[key]
	if ok {
		fmt.Println("found:", stuinf)
	} else {
		fmt.Println("not found")
	}
}
