package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const winW, winH int = 1200, 700

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Init SDL:", err)
	}

	window, err := sdl.CreateWindow("SpaceGo", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(winW), int32(winH), sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("Init window: ", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Init renderer: ", err)
		return
	}
	defer renderer.Destroy()

	firstPlayer, err := newPlayer(renderer)
	if err != nil {
		fmt.Println("Create new Player:", err)
	}

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(0, 0, 0, 0)
		renderer.Clear()

		firstPlayer.draw(renderer)
		firstPlayer.update()
		renderer.Present()
	}

}
