package main

import "fmt"

func main() {
	var scoreSlice = make([]int, 0, 5)      //定义成绩切片，元素个数为0，容量为5
	scoreSlice = Input(scoreSlice)          //录入
	fmt.Println("before sort:", scoreSlice) //显示录入的数据
	Sort(scoreSlice)                        //排序
	fmt.Println("after sort:", scoreSlice)  //输出排序后的结果
}
func Input(scoresl []int) []int {
	n := 0    //学生数量，初始为0
	temp := 0 //存储成绩的临时变量，辅助append用
	fmt.Println("input the num of students")
	fmt.Scanln(&n) //录入元素的个数
	for i := 0; i < n; i++ {
		fmt.Printf("input a int score to scoresl[%v]\n", i)
		fmt.Scanln(&temp)               //录入当前成绩到给temp
		scoresl = append(scoresl, temp) //将在切片的最后追加元素temp
	}
	return scoresl //当元素个数超过切片容量的时候，就会空间再分配，一定要返回新地址
}
func Sort(scoresl []int) {
	n := len(scoresl)          //计算切片元素的个数
	for i := 0; i < n-1; i++ { //比较n-1轮
		for j := 0; j < n-i-1; j++ { //每一轮比较n-i-1次
			if scoresl[j] < scoresl[j+1] {
				score[j], score[j+1] = score[j+1], score[j]
			}
		}
	}
}
