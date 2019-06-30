package main

import (
	"fmt"
	"net"
)

func main() {
	ips, err := net.LookupIP("www.hitwh.edu.cn")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ips)
}
