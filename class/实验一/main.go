package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Stu struct {
	stuNum  string
	stuName string
	goScore int
}

func (s *Stu) SetstuNum(stuN string) { //学号修改
	s.stuNum = stuN
}

func (s *Stu) SetstuName(stuN string) { //姓名修改
	s.stuName = stuN
}

func (s *Stu) SetstuScore(score int) { //成绩修改
	s.goScore = score
	//fmt.Println(s.goScore)
}
func (s *Stu) GetstuNum() string {
	return s.stuNum
}
func (s *Stu) GetstuScore() int {
	return s.goScore
}
func (s *Stu) GetstuName() string {
	return s.stuName
}

func menu() {
	fmt.Println("1------录入成绩")
	fmt.Println("2------查找成绩")
	fmt.Println("3------修改成绩")
	fmt.Println("4------删除成绩")
	fmt.Println("5------成绩排序")
	fmt.Println("6------成绩统计")
	fmt.Println("7------退出")
}

func Clsscr() {
	cmd := exec.Command("cmd", "/c", "cls") // /c执行字符串指定的命令然后终止
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func Goon() {
	fmt.Println("按回车键继续")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func input(sli *[]Stu) {
	fmt.Println("欢迎进入输入系统")
	var num, name string
	var sc int
	for i := 0; i < 2; i++ {
		t := Stu{}
		fmt.Print("请输入第", i+1, "个人的学号:")
		//fmt.Scanln()
		_, _ = fmt.Scanln(&num)
		//fmt.Println(num)
		t.SetstuNum(num)
		fmt.Print("请输入第", i+1, "个人的姓名")
		//_,_=fmt.Scanln()
		fmt.Scanln(&name)
		t.SetstuName(name)
		fmt.Print("请输入第", i+1, "个人的成绩")
		_, _ = fmt.Scanln(&sc)
		//fmt.Println(sc)
		t.SetstuScore(sc)
		*sli = append(*sli, t)
	}
}

func main() {
	stuslice := make([]Stu, 0, 5) //元素个数为0，容量为5
	input(&stuslice)
	for i := 0; i < 2; i++ {
		fmt.Println(i, ":")
		fmt.Println(stuslice[i].GetstuNum())
		fmt.Println(stuslice[i].GetstuName())
		fmt.Println(stuslice[i].GetstuScore())

	}
	//sort.Slice(stuslice, func(i, j int) bool { return stuslice[i].goScore < stuslice[j].goScore }) //排序
}
