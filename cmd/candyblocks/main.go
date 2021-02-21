package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// TODO: Extract controllers to its own module
// const joystickDeadZone = 8000

func main() {
	fmt.Println("Starting...")
	// TODO: Init only the needed subsystems
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	//var gameController *sdl.Joystick
	if sdl.NumJoysticks() < 1 {
		fmt.Println("Warning: No joysticks connected!")
	} else {
		fmt.Println("Joystick detected")
		fmt.Println(sdl.NumJoysticks())
		// Why extra controller ?...
		gameController := sdl.JoystickOpen(0)
		// TODO: maybe open and if nil do not use to avoid if else ...
		fmt.Println(gameController.Name())
		defer gameController.Close()
	}

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600,
		sdl.WINDOW_SHOWN)

	if err != nil {
		panic(err)
	}
	defer func() {
		panicOnErr(window.Destroy())
	}()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	panicOnErr(surface.FillRect(nil, 0))

	rect := sdl.Rect{X: 0, Y: 0, W: 100, H: 100}
	const colorRed = 0xffff0000
	panicOnErr(surface.FillRect(&rect, colorRed))
	panicOnErr(window.UpdateSurface())

	running := true
	for running {
		// TODO: Fix time step fps.
		time.Sleep(1 * time.Millisecond)
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			// This will process all the events, TODO: should it go into the main loop ?
			fmt.Println("Event", event)
			switch event := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				fmt.Println("Key event detected")
				if event.Type == sdl.KEYUP {
					if event.Keysym.Scancode == sdl.SCANCODE_RETURN || event.Keysym.Sym == sdl.SCANCODE_SPACE {
						fmt.Println("Keyup", event.Keysym.Scancode, event.Keysym.Mod, event.Keysym.Sym)
						fmt.Println("Start Game")
					} else if event.Keysym.Scancode == sdl.SCANCODE_ESCAPE {
						running = false
						fmt.Println("Keyup", event.Keysym.Scancode, event.Keysym.Mod, event.Keysym.Sym)
						fmt.Println("Escape...")
					}
				}
				fmt.Println(event.Type)
			case *sdl.JoyAxisEvent:
				fmt.Println("value:", event.Value)
			}
		}
	}
	fmt.Println("Quit")
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
