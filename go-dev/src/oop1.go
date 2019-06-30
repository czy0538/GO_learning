package main

import "fmt"

type Student struct { //定义学生成绩类型
	name  string //姓名
	stuno int    //学号
	score int    //成绩
}

func (s *Student) Stuin() {
	fmt.Println("input name of student ")
	fmt.Scanln(&s.name) //录入姓名
	fmt.Println("input stuno of student ")
	fmt.Scanln(&s.stuno) //录入学号
	fmt.Println("input score of student ")
	fmt.Scanln(&s.score) //录入成绩
}
func (s *Student) Stuout() {
	fmt.Println(*s)
}
func main() {
	stu := &Student{}
	stu.Stuin()
	stu.Stuout()
}
