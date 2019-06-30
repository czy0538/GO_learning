package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
)

var stus = make([]Stu, 0, 1000)

type Stu struct {
	stuNum   string
	stuName  string
	stuScore int
}

func (s *Stu) SetstuNum() {
	fmt.Print("请输入学生学号：")
	fmt.Scan(&s.stuNum)
}
func (s *Stu) GetstuNum() string {
	return s.stuNum
}
func (s *Stu) SetstuName() {
	fmt.Print("请输入学生姓名：")
	fmt.Scan(&s.stuName)
}
func (s *Stu) GetstuName() string {
	return s.stuName
}
func (s *Stu) SetstuScore() {
	fmt.Print("请输入学生成绩：")
	fmt.Scan(&(*s).stuScore)
	fmt.Println((*s).stuScore)
	for s.stuScore > 100 || s.stuScore < 0 {
		fmt.Print("成绩输入错误,请重新输入：")
		fmt.Scan(&s.stuScore)
	}
}
func (s *Stu) GetstuScore() int {
	return s.stuScore
}

func SetStuInfo() {
	n := 0
	fmt.Print("输入学生数量:")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var s Stu
		s.SetstuNum()
		s.SetstuName()
		s.SetstuScore()
		stus = append(stus, s)
	}
	ShowStuInfo()
}
func FindStuInfo() {
	var name string
	for name == "" {
		fmt.Print("请输入要查找学生的姓名：")
		fmt.Scan(&name)
	}
	for _, x := range stus {
		if x.stuName == name {
			fmt.Println("该学生成绩：", x.stuScore)
			return
		}
	}
	fmt.Print("没有这个学生")

}
func ShowStuInfo() {
	for i, x := range stus {
		fmt.Println(i+1, x.stuNum, x.stuName, x.stuScore)
	}
}
func ChangeStuInfo() {
	var name string
	for name == "" {
		fmt.Print("请输入要修改学生的姓名：")
		fmt.Scan(&name)
	}
	for i := 0; i < len(stus); i++ {
		if stus[i].stuName == name {
			stus[i].SetstuScore()
		}
	}
}
func DelateStuInfo() {
	var name string
	for name == "" {
		fmt.Print("请输入要删除学生的姓名：")
		fmt.Scan(&name)
	}
	for i := 0; i < len(stus); i++ {
		if stus[i].stuName == name {
			stus = append(stus[0:i], stus[i+1:]...)
			i--
		}
	}
}
func SortStuScore() {
	sort.Slice(stus, func(i, j int) bool { return stus[i].stuScore > stus[j].stuScore })
	ShowStuInfo()
}
func StatisStuScore() {
	a, b, c, d, e := 0, 0, 0, 0, 0
	sum := 0
	for _, x := range stus {
		i := x.stuScore / 10
		switch i {
		case 10:
			a++
		case 9:
			a++
		case 8:
			b++
		case 7:
			c++
		case 6:
			d++
		default:
			e++
		}
		sum += x.stuScore

	}
	fmt.Println("\nA:", a, "\nB:", b, "\nC:", c, "\nD:", d, "\nE:", e, "\nAverage:", sum*1.0/len(stus))
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
func menu() bool {
	fmt.Print("1------录入成绩  ")
	fmt.Print("2------查找成绩  ")
	fmt.Println("3------修改成绩")
	fmt.Print("4------删除成绩  ")
	fmt.Print("5------成绩排序  ")
	fmt.Println("6------成绩统计")
	fmt.Print("7------显示成绩  ")
	fmt.Print("8------清空数据  ")
	fmt.Println("9------退出系统")
	x := 0
	fmt.Print("请输入选择的操作：")
	fmt.Scanln(&x)
	switch x {
	case 1:
		SetStuInfo()
	case 2:
		FindStuInfo()
	case 3:
		ChangeStuInfo()
	case 4:
		DelateStuInfo()
	case 5:
		SortStuScore()
	case 6:
		StatisStuScore()
	case 7:
		ShowStuInfo()
	case 8:
		Clsscr()
	case 9:
		return false
	default:
		fmt.Println("选择错误\n")
	}
	return true
}
func main() {
	key := true
	for key {
		key = menu()
		Goon()
	}
}
