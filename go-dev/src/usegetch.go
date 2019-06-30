package main

import "fmt"
import termbox "github.com/nsf/termbox-go"

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

Loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				fmt.Println("You press Esc")
			case termbox.KeyF1:
				fmt.Println("You press F1")
			case termbox.KeyArrowUp:
				fmt.Println("You press 上箭头")
			default:
				break Loop
			}
		}
	}
}
