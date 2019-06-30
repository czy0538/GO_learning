package main

import "fmt"

func main() {
	const n int = 30         //成绩数组的长度
	var score [n]int         //定义成绩数组，数组长度为常量n的值
	for i := 0; i < n; i++ { //录入数据
		fmt.Printf("input a int score to score[%v]\n", i)
		fmt.Scanln(&score[i])
	}
	fmt.Println("before sort:", score) //显示录入的数据
	for i := 0; i < n-1; i++ {         //比较n-1轮
		for j := 0; j < n-i-1; j++ { //每一轮比较n-i-1次
			if score[j] < score[j+1] {
				score[j], score[j+1] = score[j+1], score[j]
			}
		}
	}
	fmt.Println("after sort:", score) //输出排序后的结果
}
