package main

import "os"

func main() {
	os.Stdout.WriteString("hello, world\n") //输出到显示器
	f, _ := os.OpenFile("test", os.O_CREATE|os.O_WRONLY, 0)
	defer f.Close()
	f.WriteString("hello, world in a file\n") //写入到test文件
}
