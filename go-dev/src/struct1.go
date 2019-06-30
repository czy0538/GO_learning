package main

import "fmt"

type Person struct { //定义了一个名为Person的结构体类型
	Name string
	Age  uint
	Sex  byte
}

func main() {
	F1 := Person{"言承旭", 37, 'm'}
	F2 := Person{"周渝民", 33, 'm'}
	F3 := Person{"吴建豪", 36, 'm'}
	F4 := Person{"朱孝天", 35, 'm'}
	fmt.Println(F1, F2, F3, F4)
}

/**Stringer是一个可用字符描述自己的类型
是一个普遍存在的接口，定义在fmt包中：
type Stringer struct {
    String() string
}
*/
func (p Person) String() string { //具体实现fmt包中的Stringer接口
	/*Sprintf 将参数列表a填写到格式控制串format的占位符中*/
	/*func Sprintf(format string, a ...interface{}) string*/
	return fmt.Sprintf("\n%v(%c,%v years)\n", p.Name, p.Sex, p.Age)
}
