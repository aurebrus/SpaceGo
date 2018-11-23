package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	backX = 1400
	backY = 800
)

type background struct {
	texture *sdl.Texture
	x, y    float64
}

func newBackground(renderer *sdl.Renderer, x, y float64) (back background) {
	back.texture = textureFromBMP(renderer, "sprites/background.bmp")

	back.x = x
	back.y = y
	return back
}

func (back *background) draw(renderer *sdl.Renderer) {
	x := back.x - backX/2
	y := back.y - backY/2

	renderer.Copy(back.texture,
		&sdl.Rect{X: 0, Y: 0, W: backX, H: backY},
		&sdl.Rect{X: int32(x), Y: int32(y), W: backX, H: backY})
}

func (back *background) update() {
	backGSpeed := 0.2
	back.y += backGSpeed

	if back.y > 800 {
		back.y = 2 * (backGSpeed)
	} else if back.y < 0 {
		back.y = 3 * (-backGSpeed)
	}

}
