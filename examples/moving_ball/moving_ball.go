package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

var window_x_size int = 800
var window_y_size int = 400

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	var event sdl.Event

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		window_x_size, window_y_size, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	running := true
	var x_pos int32 = 0
	var y_pos int32 = 0

	for running {
		drawBackground(surface)
		drawRect(surface, x_pos, y_pos)
		window.UpdateSurface()
		x_pos += 1
		y_pos += 1
		time.Sleep(10 * time.Millisecond)

		//Events catcher
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyUpEvent:
				fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
					t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
			}
		}
	}
	sdl.Quit()
}

func drawBackground(s *sdl.Surface) {
	bg := sdl.Rect{0, 0, int32(window_x_size), int32(window_y_size)}
	s.FillRect(&bg, 0xffffffff)
}

func drawRect(s *sdl.Surface, x int32, y int32) {
	rect := sdl.Rect{x, y, 50, 50}
	s.FillRect(&rect, 0xffff0000)
}
