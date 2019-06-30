package mp

import "fmt"

type Player interface { //定义播放接口
	Play(source string)
}

func Play(source, mtype string) { //定义播放函数
	var p Player //定义一个接口变量p
	switch mtype {
	case "MP3":
		p = &MP3Player{} //接口赋值，将MP3对象接入接口对象p
	case "WAV":
		//p = &WAVPlayer//方便以后扩展，暂时没有实现
	default:
		fmt.Println("Unsupported music type", mtype)
		return
	}
	p.Play(source) //调用p对应的类型所实现的paly接口
}
