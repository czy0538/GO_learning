package main

import (
	"fmt"
	"net"
)

func main()  {
	ip,err:=net.ResolveTCPAddr("tcp","www.hitwh.edu.cn:443")
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(ip)
}
