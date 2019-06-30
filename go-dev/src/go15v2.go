package main

import "fmt"
import "runtime"

//定义vector类型 实际为[]int64类型的切片,存放计算的数据
type vector []int64

var result int64 //全局变量result存放计算结果
//计算从v[i]到v[n]的各元素和
func (v vector) Dosome(i, n int, c chan int64) {
	for ; i < n-1; i++ {
		v[i+1] += v[i] //累加求和
	}
	//发送信号告诉主Goroutine已经计算完成，c的值为计算结果
	c <- v[i]
}
func (v vector) DoAll() { //计算总控模块
	var NCPU int = runtime.NumCPU() //获取CPU的逻辑核数
	//用于接收每个CPU任务完成的信号，缓冲数量为CPU的逻辑核数
	c := make(chan int64, NCPU-1)
	//按核数分解计算任务，然后开启独立Goroutine去并发计算
	for i := 0; i < NCPU-1; i++ {
		go v.Dosome(i*len(v)/(NCPU-1), (i+1)*len(v)/(NCPU-1), c)
	}
	for i := 0; i < (NCPU - 1); i++ { //等待所有的CPU任务的完成
		//获取c中的计算结果，同时也表示一个CPU计算完成了
		value := <-c
		result += value //将得到的值汇总到result
	}
}
func (v vector) DoSum() { //非并发
	for i := 0; i <= 100000000; i++ {
		result += int64(i)
	}
}
func main() {
	//设置最大使用的核心数
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.NumCPU())
	//建立切片
	v := make(vector, 100000000, 100000000)
	for i := 0; i < 100000000; i++ { //初始化
		v[i] = int64(i + 1) //每个元素的值设为i+1
	}
	v.DoAll() //调用总控模块
	fmt.Println("The result is:", result)
	result = 0
	v.DoSum()
	fmt.Println("The result is:", result)
}
