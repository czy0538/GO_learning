package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(addr)
}
