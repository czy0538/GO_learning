package main

import "fmt"

type Student struct { //定义学生成绩类型
	name  string //姓名
	stuno int    //学号
	score int    //成绩
}

func main() {
	studentmap := make(map[int]Student) //建立studentmapmap，关键字为int类型
	stu := Student{}                    //stu接收录入的学生信息，装填进studentmap
	stunum := 0                         //要查找的学号
	n := 0                              //学生数量
	fmt.Println("please input the number of students")
	fmt.Scanln(&n)           //录入学生数量
	for i := 0; i < n; i++ { //装填studentmapmap
		fmt.Println("input name of student ", i+1)
		fmt.Scanln(&stu.name) //录入姓名
		fmt.Println("input stuno of student ", i+1)
		fmt.Scanln(&stu.stuno) //录入学号
		fmt.Println("input score of student ", i+1)
		fmt.Scanln(&stu.score)      //录入成绩
		studentmap[stu.stuno] = stu //以学号为关键字装填stu到studentmap
	}
	fmt.Println("input the stuno for searching")
	fmt.Scanln(&stunum)               //输入查找的学号
	stuinfo, ok := studentmap[stunum] //查找，stuinfo代表找到的值,ok代表查找状态
	if ok {                           //找到
		fmt.Println("Find success,the info is ", stuinfo)
	} else { //没找到
		fmt.Println("Find fail")
	}
}
