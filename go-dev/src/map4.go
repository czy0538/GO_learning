package main

import "fmt"
import "os"

type Student struct { //定义学生成绩类型
	name  string //姓名
	stuno int    //学号
	score int    //成绩
}

func main() {
	studentmap := make(map[int]Student) //建立studentmap，关键字为int类型
	menu(studentmap)
}

func Input(stumap map[int]Student) { //录入数据，以map作为参数
	stu := Student{} //stu接收录入的学生信息，装填进stumap
	n := 0           //学生数量
	fmt.Println("please input the number of students")
	fmt.Scanln(&n)           //录入学生数量
	for i := 0; i < n; i++ { //装填studentmapmap
		fmt.Println("input name of student ", i+1)
		fmt.Scanln(&stu.name) //录入姓名
		fmt.Println("input stuno of student ", i+1)
		fmt.Scanln(&stu.stuno) //录入学号
		fmt.Println("input score of student ", i+1)
		fmt.Scanln(&stu.score)  //录入成绩
		stumap[stu.stuno] = stu //以学号为关键字装填stu到studentmap
	}
}
func Search(stumap map[int]Student) { //查找方法
	stunum := 0 //要查找的学号
	fmt.Println("input the stuno for searching")
	fmt.Scanln(&stunum)           //输入查找的学号
	stuinfo, ok := stumap[stunum] //查找，stuinfo代表找到的值,ok代表查找状态
	if ok {                       //找到
		fmt.Println("Find success,the info is ", stuinfo)
	} else { //没找到
		fmt.Printf("Find fail\n")
	}
}
func Display(stumap map[int]Student) { //显示
	for _, v := range stumap {
		fmt.Println(v)
	}
}
func menu(stumap map[int]Student) {
	for {
		fmt.Println("1------录入数据")
		fmt.Println("2------按学号查找")
		fmt.Println("3------显示全部")
		fmt.Println("4------退出")
		var s int
		fmt.Scanln(&s)
		switch {
		case s == 1:
			Input(stumap) //录入
		case s == 2:
			Search(stumap) //查找
		case s == 3:
			Display(stumap) //显示全部
		case s == 4:
			os.Exit(0) //退出
		default:
		}
	}
}
