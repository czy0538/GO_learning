package main

import "fmt"

func main() {
	fmt.Println("Starting the program")
	panic("A severe error occurred: stopping the program!")
	fmt.Println("Ending the program")
}

var user = os.Getenv(“USER”)
func check() {//可以在导入包的 init() 函数中检查这些
    if user == “” {
        panic(“Unknown user: no value for $USER”)
    }
}
//当发生错误必须中止程序时，panic 可以用于错误处理模式
if err != nil {
    panic(“ERROR occurred:” + err.Error())
}