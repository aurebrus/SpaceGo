package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const winW, winH = 1400, 800
const shotADelay = time.Millisecond * 7000

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

	var backgrounds []background
	for i := 0; i < 1; i++ {
		for j := 0; j < 3; j++ {
			y := (float64(j)/2)*winH + (300)
			x := float64(i) + winH
			background := newBackground(renderer, x, y)
			if err != nil {
				fmt.Println("Create new background", err)
				return
			}
			backgrounds = append(backgrounds, background)
		}
	}

	firstPlayer := newPlayer(renderer)
	if err != nil {
		fmt.Println("Create new Player:", err)
		return
	}

	var aliens []alien
	for i := 0; i < 6; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i)/5)*winH + (alienSize * 3)
			y := float64(j)*alienSize + alienSize
			alien := newAlien(renderer, x, y)
			if err != nil {
				fmt.Println("Create new Alien", err)
				return
			}
			aliens = append(aliens, alien)
		}
	}

	initTorpedoPool(renderer)
	initATorpedoPool(renderer)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(0, 0, 0, 0)
		renderer.Clear()
		for _, background := range backgrounds {
			background.draw(renderer)
		}
		for i := 0; i < 3; i++ {
			backgrounds[i].update()
		}
		firstPlayer.draw(renderer)
		firstPlayer.update()
		for _, alien := range aliens {
			alien.draw(renderer)
		}
		for i := 0; i < 18; i++ {
			aliens[i].update()
			if time.Since(aliens[i].lastShot) >= shotADelay {
				aliens[i].lastShot = time.Now()
				aliens[i].torpedoShoot()
			}

		}
		for _, tor := range torpedoPool {
			tor.draw(renderer)
			tor.update()
		}
		for _, torA := range torpedoAPool {
			torA.draw(renderer)
			torA.update()
		}
		renderer.Present()
	}
}

func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Errorf("loading renderer %v: %v", filename, err))
	}
	defer img.Free()

	texture, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("loading texture %v: %v", filename, err))
	}
	return texture
}
