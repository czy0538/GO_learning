package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Clsscr() {
	cmd := exec.Command("cmd", "/c", "cls")
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

type Stu struct {
	stuNum  string
	stuName string
	goScore int
}

func main() {

	//stuslice := make([]Stu, 0, 5)
	receive()
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
func receive() {
	stuslice := make([]Stu, 0, 5)
	var n int
	for {
		menu()
		fmt.Println("请输入您的选择：")
		fmt.Scanln(&n)
		if n == 1 {
			stuslice = input(stuslice)
			Goon()
		}
		if n == 2 {
			find(stuslice)
			Goon()
		}
		if n == 3 {
			change(stuslice)
			Goon()
		}
		if n == 4 {
			delete(stuslice)
			Goon()
		}
		if n == 5 {
			sort(stuslice)
			Goon()
		}
		if n == 6 {
			tongji(stuslice)
			Goon()
		}
		if n == 7 {
			os.Exit(0)
			Goon()
		}

	}

}
func input(x []Stu) []Stu {
	var n int
	var name string
	var card string
	var score int
	var temp Stu
	fmt.Println("请输入录入的学生数量：")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Println("请输入姓名:")
		fmt.Scanln(&name)
		fmt.Println("请输入学号:")
		fmt.Scanln(&card)
		fmt.Println("请输入go语言成绩:")
		fmt.Scanln(&score)
		temp.stuNum = card
		temp.stuName = name
		temp.goScore = score
		x = append(x, temp)
	}
	return x
}
func find(x []Stu) {
	var name string
	fmt.Println("请输入姓名:")
	fmt.Scanln(&name)
	sum := len(x)
	fmt.Println(sum) //测试用
	for i := 0; i < sum; i++ {
		if x[i].stuName == name {
			fmt.Println("姓名", x[i].stuName)
			fmt.Println("学号", x[i].stuNum)
			fmt.Println("成绩", x[i].goScore)
			break
		} else {
			fmt.Println("没找到")
		}
	}

}
func change(x []Stu) {
	var name string
	var new_score int
	fmt.Println("请输入姓名:")
	fmt.Scanln(&name)
	fmt.Println("请输入新成绩:")
	fmt.Scanln(&new_score)
	sum := len(x)
	for i := 0; i < sum; i++ {
		if x[i].stuName == name {
			x[i].goScore = new_score
			break
		}
	}
}
func delete(x []Stu) {
	var name string
	fmt.Println("请输入姓名:")
	fmt.Scanln(&name)
	sum := len(x)
	for i := 0; i < sum; i++ {
		if x[i].stuName == name {
			x[i].goScore = 0
			fmt.Println("删除成功")
			break
		}
	}
}

func sort(x []Stu) []Stu {
	n := len(x)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if x[i].goScore < x[j].goScore {
				x[i], x[j] = x[j], x[i]
			}
		}
	}
	fmt.Println("排序成功！")
	fmt.Println(x)
	return x

}
func tongji(x []Stu) {
	var a, b, c, d, e int //用来存储A,B,C,D,E五个等级的人数
	sum := len(x)
	for i := 0; i < sum; i++ {
		if x[i].goScore >= 90 && x[i].goScore <= 100 {
			a++
		}
		if x[i].goScore >= 80 && x[i].goScore < 90 {
			b++
		}
		if x[i].goScore >= 70 && x[i].goScore < 80 {
			c++
		}
		if x[i].goScore >= 60 && x[i].goScore < 70 {
			d++
		}
		if x[i].goScore >= 0 && x[i].goScore < 60 {
			e++
		}
	}
	fmt.Printf("A人数:\n", a)
	fmt.Printf("B人数:\n", b)
	fmt.Printf("C人数:\n", c)
	fmt.Printf("D人数:\n", d)
	fmt.Printf("E人数:\n", e)

}
