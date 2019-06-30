package mp

import (
	"fmt"
	"time"
)

type MP3Player struct { //mp3类型定义
	stat     int
	progress int
}

func (p *MP3Player) Play(source string) { //play接口实现
	fmt.Println("Playing MP3 music", source)
	p.progress = 0
	for p.progress < 100 {
		time.Sleep(5 * time.Second) //假装正在播放
		fmt.Printf(".")
		p.progress += 10
	}
	fmt.Println("\nFinished playing", source)
}
