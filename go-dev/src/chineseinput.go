package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readIn := bufio.NewReader(os.Stdin)
	fmt.Println("请输入姓名：")
	input, err := readIn.ReadString('\n')
	if err == nil {
		fmt.Println(input)
	}
	fmt.Println("请输入中文")
	fmt.Scanln(&input)
	fmt.Println(input)
}
