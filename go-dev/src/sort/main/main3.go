package main

import "fmt"

func main() {
	const n int = 30                   //成绩数组的长度
	var score [n]int                   //定义成绩数组，数组长度为常量n的值
	score = Input(score)               //以score作为参数，结果再返回给score
	fmt.Println("before sort:", score) //显示录入的数据
	score = sort(score)                //以score作为参数，结果再返回给score
	fmt.Println("after sort:", score)  //输出排序后的结果
}

//参数与返回值的长度要和主调函数中的score一致。
func Input(scorearr [30]int) [30]int {
	n := len(scorearr) //计算数组的长度
	for i := 0; i < n; i++ {
		fmt.Printf("input a int score to scorearr[%v]\n", i)
		fmt.Scanln(&scorearr[i])
	}
	return scorearr
}

//参数与返回值的长度要和主调函数中的score一致。
func sort(scorearr [30]int) [30]int {
	n := len(scorearr)         //计算数组的长度
	for i := 0; i < n-1; i++ { //比较n-1轮
		for j := 0; j < n-i-1; j++ { //每一轮比较n-i-1次
			if scorearr[j] < scorearr[j+1] {
				score[j], score[j+1] = score[j+1], score[j]
			}
		}
	}
	return scorearr
}
