package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	fmt.Println("Starting...")
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600,
		sdl.WINDOW_SHOWN)

	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	rect := sdl.Rect{X: 0, Y: 0, W: 200, H: 200}
	surface.FillRect(&rect, 0xffff0000)
	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			// This will process all the events, TODO: should it go into the main loop ?
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			case *sdl.KeyboardEvent:
				println("Key event detected")
				keyEvt := event.(*sdl.KeyboardEvent)
				if keyEvt.Type == sdl.KEYUP && (keyEvt.Keysym.Sym == 13 || keyEvt.Keysym.Sym == 32) {
					println("Keyup", keyEvt.Keysym.Scancode, keyEvt.Keysym.Mod, keyEvt.Keysym.Sym)
					println("Start Game")
				}
				println(keyEvt.Type)
				break
			}
		}
	}
}
