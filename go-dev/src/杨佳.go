package main

import (
	"fmt"
)

type Stu struct { //定义一个学生结构体
	stuNum  string
	stuName string
	goScore int
}

func Input(stumap *[]Stu) { //录入信息
	var stuN string
	var stuNa string
	var stus int
	var temp Stu
	fmt.Println("录入人数:")
	var n int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Println("输入学号:")
		fmt.Scanln(&stuN)
		temp.stuNum = stuN
		fmt.Println("输入姓名:")
		fmt.Scanln(&stuNa)
		temp.stuName = stuNa
		fmt.Println("输入go成绩:")
		fmt.Scanln(&stus)
		temp.goScore = stus
		*stumap = append(*stumap, temp)
	}
}

func Delet(stumap *[]Stu) { //删除特定姓名的学生信息
	fmt.Println("请输入删除要学生姓名:")
	var name string
	fmt.Scanln(&name)
	pos := 0
	for key, _ := range *stumap {
		if (*stumap)[key].stuName == name {
			fmt.Println((*stumap)[key])
			pos = key
			break
			/**stumap = append((*stumap)[:key], (*stumap)[key+1:]...)*/
		}
	}
	*stumap = append((*stumap)[:pos], (*stumap)[pos+1:]...)
}

func main() {
	stumap := make([]Stu, 0, 5)
	Input(&stumap)
	Delet(&stumap)
	fmt.Println(stumap)
}
