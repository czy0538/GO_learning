package main

import "fmt"

type Stu struct {
	stuname  string
	stunum   int
	stuscore float32
}

func (s *Stu) Setstunum() {
	snum := 0
	fmt.Println("请输入学号")
	fmt.Scanln(&snum)
	s.stunum = snum
}
func (s *Stu) Setstuname() {
	sname := ""
	fmt.Println("请输入姓名")
	fmt.Scanln(&sname)
	s.stuname = sname
}
func (s *Stu) Setstuscore() {
	var sscore float32
	fmt.Println("请输入成绩")
	fmt.Scanln(&sscore)
	s.stuscore = sscore
}
func main() {
	s := Stu{}
	s.Setstunum()
	s.Setstuname()
	s.Setstuscore()
	fmt.Println(s)
}
