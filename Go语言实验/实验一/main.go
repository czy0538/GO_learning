//180400501 曹志远
//go语言实验1
//2019年7月11日 21:25:31

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
)

type Stu struct {
	stuNum  string
	stuName string
	goScore int
}

var number = 5

//修改和显示相关函数
func (s *Stu) SetstuNum(stuN string) {
	s.stuNum = stuN
}

func (s *Stu) SetstuName(stuN string) {
	s.stuName = stuN
}

func (s *Stu) SetstuScore(score int) {
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

func Input(sli *[]Stu) {
	fmt.Println("欢迎进入输入系统")
	var num, name string
	var sc int
	for i := 0; i < number; i++ {
		t := Stu{}
		fmt.Print("请输入第", i+1, "个人的学号:")
		//fmt.Scanln()
		_, _ = fmt.Scanln(&num)
		//fmt.Println(num)
		t.SetstuNum(num)
		fmt.Print("请输入第", i+1, "个人的姓名")
		//_,_=fmt.Scanln()
		_, _ = fmt.Scanln(&name)
		t.SetstuName(name)
		fmt.Print("请输入第", i+1, "个人的成绩")
		for true {
			_, _ = fmt.Scanln(&sc)
			if sc > 100 || sc < 0 {
				println("请检查您的输入")
			} else {
				break
			}
		}
		//fmt.Println(sc)
		t.SetstuScore(sc)
		*sli = append(*sli, t)
	}
}
func Output_all(stuslice []Stu) {
	for i := 0; i < number; i++ {
		fmt.Println(i, ":")
		fmt.Println("学号", stuslice[i].stuNum)
		fmt.Println("姓名", stuslice[i].stuName)
		fmt.Println("成绩", stuslice[i].goScore)
	}
}

//修改模块
func Change_one(stuslice []Stu) {
	fmt.Println("查找依据:")
	fmt.Println("1:姓名")
	fmt.Println("2.学号")
	flag := 0
	_, _ = fmt.Scanln(&flag)

	switch flag {
	case 1:
		fmt.Print("输入你要查找的人的姓名：")
		var name string
		_, _ = fmt.Scanln(&name)
		var i int
		for i = 0; i < number; i++ {
			if stuslice[i].stuName == name {
				fmt.Print("输入你要查找的人的成绩：")
				var score int
				_, _ = fmt.Scanln(&score)
				stuslice[i].SetstuScore(score)
				fmt.Println("修改完成")
				break
			}
		}
		if i == number {
			fmt.Println("请检查您的输入")
		}

	case 2:
		fmt.Print("输入你要查找的人的学号：")
		var name string
		_, _ = fmt.Scanln(&name)
		var i int
		for i = 0; i < number; i++ {
			if stuslice[i].stuNum == name {
				fmt.Print("输入你要查找的人的成绩：")
				var score int
				_, _ = fmt.Scanln(&score)
				stuslice[i].SetstuScore(score)
				fmt.Println("修改完成")
				break
			}
		}
		if i == number {
			fmt.Println("请检查您的输入")
		}

	default:
		fmt.Println("请检查您的输入")

	}
}

//查找学生
func Find(stuslice []Stu) {
	num := number
	fmt.Print("输入你要查找的人的姓名：")
	var name string
	_, _ = fmt.Scanln(&name)
	var i int
	for i = 0; i < number; i++ {
		if stuslice[i].stuName == name {
			fmt.Print("找到啦!")
			fmt.Println("学号", stuslice[i].stuNum)
			fmt.Println("姓名", stuslice[i].stuName)
			fmt.Println("成绩", stuslice[i].goScore)
			break
		}
	}
	if i == num {
		fmt.Println("查无此人哦")
	}
}

//删除函数
func SliceRemove(s *[]Stu, index int) {
	*s = append((*s)[:index], (*s)[index+1:]...)
	number--
}

//删除模块
func Del(stuslice []Stu) {
	fmt.Print("输入你要查找的人的姓名：")
	var name string
	_, _ = fmt.Scanln(&name)
	var i int
	for i = 0; i < number; i++ {
		if stuslice[i].stuName == name {
			SliceRemove(&stuslice, i)
			fmt.Println("删除成功")
		}
	}
	if i == number {
		fmt.Println("查无此人哦")
	}
}

func Sort_less(stuslice []Stu) {
	sort.Slice(stuslice, func(i, j int) bool { return stuslice[i].goScore < stuslice[j].goScore })
}

func Count(stuslice []Stu) {
	class := map[string]int{}
	class["A"] = 0
	class["B"] = 0
	class["C"] = 0
	class["D"] = 0
	class["E"] = 0
	sum := 0
	var avg float64 = 0.0
	for i := 0; i < number; i++ {
		switch sc := stuslice[i].goScore; {
		case sc >= 90 && sc <= 100:
			class["A"]++
		case sc < 90 && sc >= 80:
			class["B"]++
		case sc < 80 && sc >= 70:
			class["C"]++
		case sc < 70 && sc >= 60:
			class["D"]++
		default:
			class["E"]++
		}
		sum += stuslice[i].goScore
	}
	avg = float64(sum * 1.0 / number)
	fmt.Println("平均分为:", avg)
	for i, j := range class {
		fmt.Println(i, j)
	}

}
func Menu() {
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

func main() {
	stuslice := make([]Stu, 0, number) //元素个数为0，容量为5
	var m int = 0
	for true {
	flag:
		Menu()

		_, _ = fmt.Scanln(&m)
		switch m {
		case 1:
			Input(&stuslice)
			Output_all(stuslice)
			Goon()
			Clsscr()
			goto flag

		case 2:
			Find(stuslice)
			Goon()
			Clsscr()
			goto flag
		case 3:
			Change_one(stuslice)
			Output_all(stuslice)
			Goon()
			Clsscr()
			goto flag
		case 4:
			Del(stuslice)
			Output_all(stuslice)
			Goon()
			Clsscr()
			goto flag
		case 5:
			Sort_less(stuslice)
			Output_all(stuslice)
			Goon()
			Clsscr()
			goto flag
		case 6:
			Count(stuslice)
			Goon()
			Clsscr()
			goto flag
		case 7:
			fmt.Println("谢谢您的使用")
			os.Exit(0)
		}

	}
}
