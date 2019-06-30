package main

import (
	"fmt"
	"net"
)

func main() {
	//解析服务器UDP地址
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:7070")
	if err != nil {
		fmt.Println(err)
		return
	}
	//链接服务器
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close() //关闭链接
	//向服务器发送数据
	conn.Write([]byte("Hello Server"))
	var buf [512]byte
	//读取服务器的响应信息
	n, _, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf[0:n]))
}
