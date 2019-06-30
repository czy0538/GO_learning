package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"strings"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

type Subscriber struct {
	user *websocket.Conn
	time time.Time
	name string
}

var messages []string
var users []Subscriber

func Echo(ws *websocket.Conn) {
	var User Subscriber

	//获取用户信息
	User.user = ws
	User.time = time.Now()
	n, _ := receive(ws)
	if strings.Contains(n, "n-") == true { // 判断从服务器端收到的消息是否含有”n-“
		User.name = n[2:] // 建立切片，将客户姓名取出
	}
	users = append(users, User)

	time_start := time.Now()

	for {
		online_users()
		reply, ok := receive(ws) //comma-ok模式，防止网页退出时导致程序崩溃

		if strings.Contains(reply, "*-exit-*") == true {
			ws.Close()
			for _, u := range users {
				if u.user == ws {
					fmt.Println(u.name + "退出")
				}
			}
		}

		if !ok { // 如果接受失败，删除用户
			rec_err_deal(ws)
			break
		}

		messages = append(messages, reply) //记录聊天信息
		send(reply)

		fmt.Println(reply)
		t2 := time.Since(time_start)
		if t2 > 600000 { // 每十分钟把聊天记录写入文件
			chat_data()
			time_start = time.Now()
		}
	}
}

//接受用户信息失败时删除该用户
func rec_err_deal(ws *websocket.Conn) {
	for i, u := range users {
		if u.user == ws { // 找到离线的用户
			ws.Close()
			users = append(users[:i], users[i+1:]...)
		}

		fmt.Println("用户已断开链接！")
	}
}

//显示已登录用户
func online_users() {
	fmt.Println(users)
	for _, u := range users {
		for _, us := range users {
			name := "--在线用户："
			name = name + u.name
			err := websocket.Message.Send(us.user, name)
			//fmt.Println(name)
			if err != nil {
				fmt.Println("发送在线用户名错误！", err)
			}
		}
	}
}

//接受来自网页的数据
func receive(ws *websocket.Conn) (string, bool) {
	var reply string
	timeout := make(chan bool, 1)
	rec := make(chan bool, 1)
	flag := true

	//连接超时控制
	go func() {
		time.Sleep(time.Second * 100) // 设定为100秒不发送消息则断开链接
		timeout <- true
	}()

	go func() {
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Fail to receive:", err)
			flag = false
		}
		rec <- true
	}()

	select {
	case <-rec:
		return reply, flag

	case <-timeout:
		for _, u := range users {
			if u.user == ws {
				reply = u.name + "连接超时，退出。"
				fmt.Println(reply)
			}
		}
		ws.Close()
		flag = false
		return reply, flag //comma-ok 模式，当flag为false是进行进一步处理
	}
}

//向所有用户发送数据并删除离线用户
func send(reply string) {
	var wg sync.WaitGroup

	for i, user := range users { // 遍历所有用户
		wg.Add(1)
		go func(user Subscriber, i int) { // 运用并发
			err := websocket.Message.Send(user.user, reply)
			if err != nil {
				user.user.Close()
				fmt.Println("Fail to send!", err)
				users = append(users[:i], users[i+1:]...) // 若用户已离开，则删除用户
				fmt.Println("已删除该用户！")
			}
			wg.Done()
		}(user, i)
		wg.Wait()
	}
}

//存储聊天数据
func chat_data() {
	file, err := os.Open("chat_datas.txt") // 写入文件
	if err != nil {
		fmt.Println("文件创建失败:", err)
	}
	for _, con := range messages {
		con = con + "\n"
		file.WriteString(con)
	}
}

func main() {

	http.Handle("/", websocket.Handler(Echo))

	//监听端口
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
