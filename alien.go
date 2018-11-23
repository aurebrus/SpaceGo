package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	alienSize = 90
)

type alien struct {
	texture *sdl.Texture
	x, y    float64
}

func newAlien(renderer *sdl.Renderer, x, y float64) (al alien) {
	al.texture = textureFromBMP(renderer, "sprites/aliencraft.bmp")

	al.x = x
	al.y = y

	return al
}

func (al *alien) draw(renderer *sdl.Renderer) {
	x := al.x - alienSize/2
	y := al.y - alienSize/2

	renderer.Copy(al.texture,
		&sdl.Rect{X: 0, Y: 0, W: alienSize, H: alienSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: alienSize, H: alienSize})
}

func (al *alien) update() {
	alienSpeed := 0.4
	al.x += alienSpeed

	if al.x > 1150 {
		alienSpeed = -alienSpeed
	} else if al.x < 50 {
		alienSpeed = -alienSpeed
	}

}
