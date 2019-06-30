package main

import (
	"fmt"
)

type Student struct {
	Name   string
	Stunum string
	Score  int
}

func main() {
	var keyname string
	var n int
	stu := Student{}
	studentmap := make(map[string]Student)
	fmt.Println("input the num of students")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Println("input the name of student", i+1)
		fmt.Scanln(&stu.Name)
		fmt.Println("input the stuno of student", i+1)
		fmt.Scanln(&stu.Stunum)
		fmt.Println("input the score of student", i+1)
		fmt.Scanln(&stu.Score)
		studentmap[stu.Name] = stu
	}
	fmt.Println("Input a stunum for searching")
	fmt.Scanln(&keyname)
	data, ok := studentmap[keyname]
	if !ok {
		fmt.Println("not find")
	} else {
		fmt.Println("the result is\n", data)
	}
}
