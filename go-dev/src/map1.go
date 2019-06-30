package main

import "fmt"

type Student struct { //定义学生成绩类型
	name  string //姓名
	stuno int    //学号
	score int    //成绩
}

func main() {
	//用内置函数make()创建一个新map，名称为studentmap,
	//map[string]Student表明 该map的key为string类型,存放的数据为Student类型
	//我们将这个string类型的key对应map中的姓名字段
	studentmap := make(map[string]Student)
	//装填数据即为哈希表进行元素赋值,装填函数go map已经为我们封装好,直接使用即可
	//关键字Zhao对应的pos中 放入元素Student{"Zhao", 140410101, 99}
	studentmap["Zhao"] = Student{"Zhao", 140410101, 99}
	studentmap["Qian"] = Student{"Qian", 140410102, 88}
	studentmap["Sun"] = Student{"Sun", 140410103, 77}
	studentmap["Li"] = Student{"Li", 140410104, 66}
	//查找姓名为Sun的成绩，查找的具体过程go map也已经封装好,直接使用即可
	stuinfo, ok := studentmap["Sun"] //stuinfo代表找到的值,ok代表查找成功与否
	if ok {
		fmt.Println("Find success,the info is ", stuinfo)
	} else {
		fmt.Println("Find fail")
	}
}
