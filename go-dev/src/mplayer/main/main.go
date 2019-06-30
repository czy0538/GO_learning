package main

import (
	"bufio" //输入输出缓存包
	"fmt"
	"mplayer/mlib" //这个包里面定义音乐库数据结构
	"mplayer/mp"   //这个包里面定义接口、播放函数、MP3类型及接口实现
	"os"           //系统包
	"strconv"      //字符串转换包
	"strings"      //字符串处理包
)

var lib *mlib.MusicManager //音乐库指针对象lib
var id int = 1             //音乐编号id
//音乐库管理
func handleLibCommand(tokens []string) { //参数为字符串切片
	switch tokens[1] {
	case "list": //列出音乐清单
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i) //获取索引为i的音乐
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add": //增加音乐
		if len(tokens) == 6 { //如果命令的长度为6，正确的增加命令split后的长度为6
			id++
			//音乐入库
			lib.Add(&mlib.Music{strconv.Itoa(id), tokens[2], tokens[3], tokens[4], tokens[5]})
		} else { //提示增加音乐命令的正确格式是什么
			fmt.Println("USAGE:lib add <name><artist><source><type>")
		}
	case "remove": //删除音乐
		if len(tokens) == 3 { //如果命令的长度为3，正确的删除命令split后的长度为3
			_, i := lib.Find(tokens[2]) //找要删除的音乐
			if i != -1 {                //找到
				lib.Remove(i)
			} else { //乐库为空或没找到
				fmt.Println("music", tokens[2], "is not exist in lib")
			}
		} else { //提示删除音乐命令的正确格式是什么
			fmt.Println("USAGE:lib remove<name>")
		}
	default: //提示命令无法识别
		fmt.Println("Unrecognized lib command:", tokens[1])
	}
}
func handlPlayCommand(tokens []string) { //音乐播放
	if len(tokens) != 2 { //如果命令长度不等于2，正确的播放命令split的长度为2
		fmt.Println("USAGE:play<name>") //提示正确的格式
		return
	}
	e, _ := lib.Find(tokens[1]) //查找音乐
	if e == nil {               //没找到
		fmt.Println("The music", tokens[1], "does not exist.")
		return
	}
	mp.Play(e.Source, e.Type) //执行播放函数
}

func main() {
	fmt.Println(`
			Enter following commands  to control the player:
			lib list -- View the existing music lib
			lib add <name><artist><source><type> --Add a music to the music lib
			lib remove <name> --Remove the specified music from the lib
			play <name> --Play the specified music
			`) //Println 用``可以对其中字符串按照所见即所得模式管理
	lib = mlib.NewMusicManager()   //新建立一个音乐库对象
	r := bufio.NewReader(os.Stdin) //建立一个缓存对象r，可以接收来自控制台的输入
	for {
		fmt.Print("Enter command->")
		rawLine, _, _ := r.ReadLine()   //读取一行数据到rawLine数组
		line := string(rawLine)         //将rawLine转换成字符串赋给line
		if line == "q" || line == "e" { //如果等于q或e退出程序
			break
		}
		tokens := strings.Split(line, " ") //用空格分隔命令串
		if tokens[0] == "lib" {            //如果第一部分为lib
			handleLibCommand(tokens) //执行处理音乐库管理功能
		} else if tokens[0] == "play" { //如果第一部分为play
			handlPlayCommand(tokens) //执行音乐播放功能
		} else { //否则 提示命令无法识别
			fmt.Println("Unrecognized command:", tokens[0])
		}
	}
}
