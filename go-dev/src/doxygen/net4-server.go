package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

const ( //将命令关键字定义为常量
	LS  = "LS"
	CD  = "CD"
	PWD = "PWD"
)

func main() {
	//解析tcp地址
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":7076")
	//解析出错
	checkError(err)
	//解析正确则侦听，在7076端口侦听
	listener, err1 := net.ListenTCP("tcp", tcpAddr)
	checkError(err1) //监听失败
	//监听成功，等待连接提示
	fmt.Println("等待客户端的链接")
	for {
		//等待客户端的连接
		conn, err2 := listener.Accept()
		if err != nil {
			/*通常服务端为一个服务，不会因为错误而退出。出错后，继续
			  等待下一个连接请求*/
			fmt.Println(err2)
			continue
		}
		fmt.Println("收到客户端的请求", conn.RemoteAddr())
		//开启一个并发的goroutine处理来自链接对象指令
		go ServeClient(conn)
	}
}
func ServeClient(conn net.Conn) {
	//处理完毕后关闭，优雅的defer防止最后忘记关闭
	defer conn.Close()
	str := ReadData(conn) //读取数据
	//读取命令失败，则告诉客户接收数据出错
	if str == "" {
		SendData(conn, "接收数据时出错")
		return
	}
	//读取成功
	fmt.Println("收到命令：", str)
	switch str {
	case LS: //列出当前目录的文件及文件夹清单
		ListDir(conn)
	case PWD: //给出当前目录
		Pwd(conn)
	default: //
		if str[0:2] == CD { //前两个字节为CD的情况
			Chdir(conn, str[3:])
		} else { //命令不正确
			SendData(conn, "命令错误")
		}
	}
}
func Chdir(conn net.Conn, s string) {
	err := os.Chdir(s) //改变目录到指定目录
	if err != nil {    //改变失败，发送错误信息给客户端
		SendData(conn, err.Error())
	} else { //改变成功，发送OK给客户端
		SendData(conn, "OK")
	}
}

//列出当前目录的文件及子文件夹
func ListDir(conn net.Conn) {
	//读出当期路径下的文件信息
	files, err := ioutil.ReadDir(".")
	if err != nil { //读取失败，并告知客户端
		SendData(conn, err.Error())
		return
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

/*读取数据*/
func ReadData(conn net.Conn) string {
	var data bytes.Buffer
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
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
func Pwd(conn net.Conn) {
	//获取当前工作目录
	s, err := os.Getwd()
	if err != nil { //获取失败，发送错误信息给客户端
		SendData(conn, err.Error())
	} else { //获取成功，将当前工作目录发给客户端
		SendData(conn, s)
	}
}
func checkError(err error) { //出错则退出
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
