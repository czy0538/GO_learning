package main

import "fmt"
import "math/rand"
import "time"

func main() {

	a := [100]int{} //定义数组a
	//b := a[0:10]               //定义切片b为a[0]-a[9]
	for i := 0; i < 100; i++ { //a数组赋值为1-100
		a[i] = i + 1
	}
	rand.Seed(time.Now().Unix())
	for j := 0; j < 10; j++ { //产生多少个就循环多少次
		k := j + rand.Intn(100-j)     // j + rand() % (100 - j)产生一个范围[j,100)的随机数
		a[j], a[k] = swap(a[j], a[k]) //那么将下标k的数据和以j为下标的数据swap肯定不重复
	}
	//fmt.Println(b)
	fmt.Println(a[0:10])
}
func swap(a, b int) (int, int) {
	return b, a
}
