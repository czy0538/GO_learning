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
			case termbox.KeyCtrlA:
				fmt.Println("You press Ctrl+A")
			case termbox.KeyCtrlC:
				fmt.Println("You press Ctrl+C")
			case termbox.KeyCtrlV:
				fmt.Println("You press Ctrl+V")
			default:
				break Loop
			}
		}
	}
}
