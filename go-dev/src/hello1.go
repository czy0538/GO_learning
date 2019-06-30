package main

import "fmt"
import (
	"runtime"
)

func main() {
	names := []string{"帅天强", "马晓乾", "姚雨辰", "郭逸飞", "谢莫莉"}
	for _, name := range names {
		go fmt.Println("hello ", name)
		runtime.Gosched()
	}

}
