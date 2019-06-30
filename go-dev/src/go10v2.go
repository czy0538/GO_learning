package main

import "fmt"
import "sync"

//golang中的同步是通过sync.WaitGroup来实现的
var waitgroup sync.WaitGroup

func main() {
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		//每创建一个goroutine，就把任务队列中任务的数量+1
		waitgroup.Add(1)
		go func(who string) {
			fmt.Printf("Hello,%s.\n", who)
			//任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
			waitgroup.Done()
		}(name)
	}
	waitgroup.Wait() //这里会发生阻塞,直到队列中所有的任务结束就会解除阻塞
}
