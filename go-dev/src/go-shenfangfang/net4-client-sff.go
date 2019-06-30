package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	LS   = "LS"
	CD   = "CD"
	PWD  = "PWD"
	QUIT = "QUIT"
)

func main() {
	//录入服务器的TCP地址
	tcpAddrStr := InputTCPAddr()
	//TCP地址合法性检查
	tcpAddr, err := net.ResolveTCPAddr("tcp", tcpAddrStr)
	if err != nil { //TCP地址无效，退出系统
		fmt.Println(err)
		os.Exit(0)
	}
	//服务器TCP地址有效，则链接服务器。客户机地址为nil
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil { //链接失败，退出系统
		fmt.Println(err)
		os.Exit(0)
	}
	//链接成功，conn为链接对象
	for { //循环录入并处理命令，直至退出
		cmd := InputCMD()              //录入命令
		res := ProcessCMD_C(conn, cmd) //处理命令
		//res==false 时候，退出系统
		if !res {
			fmt.Println("与服务器:", conn.RemoteAddr(), "的通信结束")
			conn.Close() //关闭连接
			break
		}
	}
}

//输入服务器的TCP地址
func InputTCPAddr() string {
	var TCPAddr string //TCP地址串
	fmt.Printf("请输入服务器的TCP地址(default: 127.0.0.1:7076):\n")
	fmt.Scanln(&TCPAddr)
	if len(TCPAddr) == 0 {
		TCPAddr = "127.0.0.1:7076"
	}
	return TCPAddr
}

//录入命令
func InputCMD() string {
	fmt.Printf("请输入命令:\n")
	//建立一个buffer reader (带缓冲的读取器)
	reader := bufio.NewReader(os.Stdin)
	//录入数据，以回车结束
	cmdline, err := reader.ReadString('\n')
	//录入失败，退出系统
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	//录入成功，去掉两端的空格
	cmdline = strings.TrimSpace(cmdline)
	//统一转换成大写字母
	cmdline = strings.ToUpper(cmdline)
	return cmdline
}

//处理命令
func ProcessCMD_C(conn net.Conn, cmd string) bool {
	var res bool = true //res返回值，默认true
	//将命令按空格分隔成2部分
	cmdarr := strings.SplitN(cmd, " ", 2)
	switch cmdarr[0] { //根据第一部分决定执行哪些命令
	case LS:
		//连接成功后，发送命令
		SendData(conn, LS)
		//读取服务器端返回的结果，并输出
		fmt.Println(ReadData(conn))
	case CD:
		if len(cmdarr) <= 1 {
			fmt.Println("命令错误！")
		} else {
			SendData(conn, CD+" "+strings.TrimSpace(cmdarr[1]))
			fmt.Println(ReadData(conn))
		}
	case PWD:
		SendData(conn, PWD)
		fmt.Println(ReadData(conn))
	case QUIT:
		SendData(conn, QUIT)
		res = false
	default:
		fmt.Println("命令错误！")
	}
	return res
}

//向服务器端发送数据
func SendData(conn net.Conn, data string) {
	buf := []byte(data)
	/*向byte 字节里添加结束标记*/
	buf = append(buf, 0)
	_, err := conn.Write(buf) //发送数据，使用conn的write方法
	if err != nil {           //发送数据出错
		fmt.Println(err)
	}
}

/*读取服务器端返回的数据*/
func ReadData(conn net.Conn) string {
	//data是个缓冲byte类型的缓冲器，这个缓冲器里存放着都是byte
	var data bytes.Buffer
	var buf [512]byte
	for {
		//接收数据到buf，使用conn.Read方法，n存储接收的字节数量
		n, err := conn.Read(buf[0:])
		if err != nil { //接收失败
			fmt.Println(err)
			return ""
		}
		//我们的数据以0做为结束的标记
		if buf[n-1] == 0 {
			//n-1去掉结束标记0
			data.Write(buf[0 : n-1])
			break
		} else {
			data.Write(buf[0:n])
		}
	}
	return string(data.Bytes())
}
