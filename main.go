package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const winW, winH int = 1200, 800

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

	img, err := sdl.LoadBMP("sprites/playercraft.bmp")
	if err != nil {
		fmt.Println("Init player: ", err)
		return
	}

	playerTexture, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		fmt.Println("Init texture:", err)
		return
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
		renderer.Copy(playerTexture,
			&sdl.Rect{X: 0, Y: 0, W: 115, H: 115},
			&sdl.Rect{X: 0, Y: 0, W: 115, H: 115})
		renderer.Present()
	}

}
