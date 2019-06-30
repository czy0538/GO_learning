package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var n int
	var f float64
	var sum float64
	//fmt.Scanln(&n)
	n = 1000
	ch := make(chan float64, 100)
	sum = 0
	for k := 0; k < n; k++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- math.Pow(-1, float64(k)) / float64(2*k+1)
		}()
		f = <-ch
		sum = sum + f
	}

	close(ch)
	wg.Wait()
	fmt.Println(sum * 4)
}

/*package main

import "fmt"

// 线程数

const N = 8

var n int

// 工作者，从输入管道中读取数字并计算出结果放到输出管道

func worker(in chan int, out chan float64) {
	for i := range in {
		if i%2 == 0 {
			i := float64(i)
			out <- 1.0 / (i*2.0 + 1.0)
		} else {
			i := float64(i)
			out <- -1.0 / (i*2.0 + 1.0)
		}
	}
}

// 发送者，向管道中填充所需计算的数字

func send(in chan int) {
	for i := 0; i <= n; i++ {
		in <- i
	}
}

// 接收者从输出管道接受计算结果，并计算出pi

func receive(out chan float64) {
	var pi float64
	for i := 0; i <= n; i++ {
		pi += <-out
	}
	pi *= 4
	fmt.Println(pi)
}

// 创建管道，读取n，创建工作者和接受者线程并调用接受者等待结果

func main() {
	in, out := make(chan int, 100), make(chan float64, 100)
	fmt.Scanf("%d", &n)
	for i := 0; i < N; i++ {
		go worker(in, out)
	}
	go send(in)
	receive(out)
}
*/
/*package main

import "fmt"
import "math"

//import "strconv"

func num(ch chan float64, cnt int) {
	for k := 0; k < cnt; k++ {

		ch <- (math.Pow(-1, float64(k)) / (2*float64(k) + 1))
		fmt.Println("写入数据：", math.Pow(-1, float64(k))/(2*float64(k)+1), "k=", k)

	}
	close(ch)
}
func sum(ch chan float64, done chan bool) {
	sum := 0.0
	for f := range ch {
		f = <-ch
		fmt.Println("收到 f=", f)
		sum = (sum + f) * 4
		fmt.Println("sum=", sum)

	}
	/*a := strconv.FormatFloat(sum, 'f', 5, 64)
	fmt.Println("pi=", a)*/
/*	fmt.Println("pi=", sum)

	done <- true
}*/
/*func main() {
n := 1000
/*fmt.Println("input n:")
fmt.Scanln(&n)*/
/*	ch := make(chan float64)
	done := make(chan bool)
	go num(ch, n)
	go sum(ch, done)
	<-done
}*/
