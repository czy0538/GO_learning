//180400501 曹志远
//2019年7月2日 15:14:34
// 作业1 两个整数的交换程序
package main

import "fmt"

func swap(a, b int) (int, int) {
	c, d := b, a
	return c, d
}
func main() {
	var a, b int
	fmt.Println("请依次输入要交换的数据")
	fmt.Scan(&a, &b)
	fmt.Println(swap(a, b))
}
