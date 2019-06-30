package main

import "fmt"
import "time"

func main() {
	ch := make(chan int, 5)
	sign := make(chan byte, 2)
	go func() {
		for i := 0; i < 5; i++ { //执行了5次发送操作，每次间隔1s
			ch <- i
			time.Sleep(1 * time.Second)
		}
		close(ch) //关闭通道
		fmt.Println("The channel is closed.")
		sign <- 0 //
	}()
	go func() {
		for {
			//接收操作，这里使用了接收的多重返回值方式接收
			//e代表值，ok是布尔值，如果是false表示关闭，true表示未关闭
			e, ok := <-ch
			fmt.Printf("%d(%v)\n", e, ok)
			if !ok { //如果ok为false表示通道已经关闭，退出循环
				break
			}
			time.Sleep(2 * time.Second) //每次接收操作间隔2s
		}
		fmt.Println("Done.") //输出
		sign <- 1
	}()
	//sign通道的作用是推迟主Goroutine被运行完成的时间
	//是两个go 匿名函数对应的 Goroutine顺利执行的关键
	//读取通道sign的值，前提是通道 sign被写入了一个值，否则该读取操作就出于阻塞等待状态
	//因为有两个go 匿名 goroutine，他们执行完毕后都要向通道sign写入数据，因此要接收两次
	<-sign
	<-sign
}
