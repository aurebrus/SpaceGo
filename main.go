package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const winW, winH int = 1200, 1000

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Init SDL:", err)
	}

	window, err := sdl.CreateWindow("SpaceGO", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(winH), int32(winH), sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("Init window: ", err)
		return
	}
	defer window.Destroy()
}
