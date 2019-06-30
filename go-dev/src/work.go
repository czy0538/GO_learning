//限制输入范围，避免超出内存限制
package main

import (
	"fmt"
	"math"
)

/*
#include <stdio.h>
*/
import "C"

type OverflowError struct { //超出范围错误
	Word int
}

//实现String接口定义错误输出格式
func (e *OverflowError) String() string {
	return fmt.Sprintf("OverflowError: error input %d,out of range.", e.Word)
}

func main() {

	chan100 := make(chan float64, 100)
	var result float64
	//=============================
AGAIN:
	count, err := Input()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Please try again.")
		goto AGAIN
	}
	//=============================
	for i := 0; i <= count; i++ {
		go routine(i, chan100)
	}

	for i := 0; i <= count; i++ {
		result += <-chan100
	}

	fmt.Println("the result of pi is : ", result*4)
}

//输入检测
func Input() (count int, Err error) {
	defer func() {
		if r := recover(); r != nil {
			Err = fmt.Errorf("Sorry, %v", r)
			count = 0
		}
	}()

	fmt.Println("你想建立多少个goroutine(1-1,000,000)?")
	fmt.Scanf("%d", &count)
	C.getchar()

	if count <= 0 || count > 1000000 {
		panic(&OverflowError{count})
	}

	return
}

func routine(i int, ch chan float64) {
	k := float64(i)
	ch <- math.Pow(-1, k) / (2*k + 1)
}
