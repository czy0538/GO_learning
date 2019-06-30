package main

import "fmt"

type Student struct { //定义学生成绩类型
	name  string //姓名
	stuno int    //学号
	score int    //成绩
}

/**Stringer是一个可用字符描述自己的类型
是一个普遍存在的接口，定义在fmt包中：
type Stringer struct {
    String() string
}
*/
func (s Student) String() string { //具体实现fmt包中的Stringer接口
	/*Sprintf 将参数列表a填写到格式控制串format的占位符中*/
	/*func Sprintf(format string, a ...interface{}) string*/
	return fmt.Sprintf("\n姓名	学号	成绩\n%v	%v %v\n", s.name, s.stuno, s.score)
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
