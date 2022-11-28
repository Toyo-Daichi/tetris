package main

import (
	"fmt"
	"time"

	termbox "github.com/nsf/termbox-go"
)

func drawLine(x int, y int, s string) {
	color := termbox.ColorDefault
	backgroudColor := termbox.ColorDefault
	runes := []rune(s)

	for i := 0; i < len(runes); i += 1 {
		termbox.SetCell(x+i, y, runes[i], color, backgroudColor)
	}
}

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawLine(0, 0, "Press ESC to exit.")
	drawLine(2, 1, fmt.Sprintf("date: %v", time.Now()))
	termbox.Flush()
}

func pollEvent() {
	draw()
	for {
		switch env := termbox.PollEvent(); env.Type {
		case termbox.EventKey:
			switch env.Key {
			case termbox.KeyEsc:
				return
			default:
				draw()
			}
		default:
			draw()
		}
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	pollEvent()
}
