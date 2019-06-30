package main

import (
	"fmt"
	"net"
)

func main() {
	//解析udp地址
	addr, err := net.ResolveUDPAddr("udp", ":7070")
	if err != nil {
		fmt.Println(err)
		return
	}
	//在7070端口监听
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	for { //循环接收数据，处理数据
		var buf [1024]byte
		n, addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		//开启新线程处理客户端的数据
		go HandleClient(conn, buf[0:n], addr)
	}
}

//处理客户端数据
func HandleClient(conn *net.UDPConn, data []byte, addr *net.UDPAddr) {
	fmt.Println("收到数据：" + string(data))
	conn.WriteToUDP([]byte("OK,数据已到"), addr)
}
