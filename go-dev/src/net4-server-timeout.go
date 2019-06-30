package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
)

const ( //将命令关键字定义为常量
	LS   = "LS"
	CD   = "CD"
	PWD  = "PWD"
	QUIT = "QUIT"
)

func main() {
	//解析tcp地址
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":7076")
	if err != nil { //解析出错
		fmt.Println(err)
		os.Exit(0)
	}
	//解析正确则侦听，在7076端口侦听
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil { //监听失败
		fmt.Println(err)
		os.Exit(0)
	}
	//监听成功，等待连接提示
	fmt.Println("等待客户端的链接")
	for {
		//接收客户端的链接
		conn, err := listener.Accept()
		if err != nil { //接收失败
			/*通常服务端为一个服务，不会因为错误而退出。
			出错后，继续等待下一个链接请求*/
			fmt.Println(err)
			continue
		}
		//成功接收链接，并显示远程客户端的地址
		fmt.Println("收到链接，自客户端:", conn.RemoteAddr())
		//开启一个Goroutine处理来自客户指令
		go ProcessClient(conn)
	}
}

//处理来自客户端的指令
func ProcessClient(conn net.Conn) {
	for {

		cmdstr := ReceiveCMD(conn) //接收命令/获取命令
		if cmdstr == "" {
			break //超时或接收命令异常退出
		}
		res := ExecCMD_S(conn, cmdstr) //执行命令
		if !res {
			fmt.Println("与客户端:", conn.RemoteAddr(), "的通信结束")
			conn.Close() //关闭与客户端的链接
			break
		}
	}
}

//接收命令/获取命令
func ReceiveCMD(conn net.Conn) string {
	timeout := make(chan bool, 1)   //超时限制通道
	cmdchan := make(chan string, 1) //命令通道
	go func() {                     //超时Goroutine
		time.Sleep(10 * time.Second)
		timeout <- true
	}()
	go func() {
		cmdstr := ReadData(conn) //读取命令/接收数据
		//读取命令失败，则告诉客户接收数据出错
		if len(cmdstr) == 0 {
			fmt.Println("与客户端:", conn.RemoteAddr(), "的通信结束")
			return
		}
		//读取成功
		fmt.Println("收到命令：", cmdstr)
		cmdchan <- cmdstr
	}() //成功接收命令
	select {
	case value := <-cmdchan:
		if len(value) != 0 {
			return value
		} else {
			conn.Close()
			return ""
		}
	case <-timeout:
		fmt.Println("由于超时与客户端:", conn.RemoteAddr(), "的通信结束")
		SendData(conn, "获取命令失败，通信结束")
		conn.Close() //关闭链接
		return ""
	}
}

//执行命令
func ExecCMD_S(conn net.Conn, cmdstr string) bool {
	var res bool = true //res返回值，默认true
	switch cmdstr {
	case "":
		res = false
	case LS: //列出当前目录的文件及文件夹清单
		ListDir(conn)
	case PWD: //给出当前目录
		Pwd(conn)
	case QUIT: //客户端结束链接指令
		res = false
	default: //其他情况
		if cmdstr[0:2] == CD { //前两个字节为CD的情况
			Chdir(conn, cmdstr[3:])
		} else { //命令不正确
			SendData(conn, "命令错误")
		}
	}
	return res
}

//LS命令：列出当前路径下的文件及子文件夹
func ListDir(conn net.Conn) {
	//读出当前路径下的文件信息
	files, err := ioutil.ReadDir(".")
	if err != nil { //读取失败，并告知客户端
		SendData(conn, err.Error())
	}
	//读取成功
	var str string
	for i, j := 0, len(files); i < j; i++ {
		f := files[i]
		str += f.Name() + "\t"
		if f.IsDir() { //如果是文件夹
			str += "dir\r\n"
		} else { //如果是文件
			str += "file\r\n"
		}
	}
	//发送给客户端文件信息及子文件夹信息
	SendData(conn, str)
}

//CD命令：改变当前路径
func Chdir(conn net.Conn, s string) {
	err := os.Chdir(s) //改变目录到指定目录
	if err != nil {    //改变失败，发送错误信息给客户端
		SendData(conn, err.Error())
	} else { //改变成功，发送OK给客户端
		SendData(conn, "OK")
	}
}

//PWD命令：获取服务器当前路径
func Pwd(conn net.Conn) {
	//获取当前工作目录
	s, err := os.Getwd()
	if err != nil { //获取失败，发送错误信息给客户端
		SendData(conn, err.Error())
	} else { //获取成功，将当前工作目录发给客户端
		SendData(conn, s)
	}
}

//从客户端读取数据
func ReadData(conn net.Conn) string {
	var data bytes.Buffer
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:]) //n为字节数
		if err != nil {
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

//发送数据给客户端
func SendData(conn net.Conn, data string) {
	buf := []byte(data)
	/*向byte 字节里添加结束标记*/
	buf = append(buf, 0)
	_, err := conn.Write(buf)
	if err != nil {
		fmt.Println(err)
	}
}
