package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.InterfaceAddrs()
	fmt.Println(addr)
}
