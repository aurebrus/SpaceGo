package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	alienSize = 90
)

type alien struct {
	texture *sdl.Texture
	x, y    float64
}

func newAlien(renderer *sdl.Renderer, x, y float64) (al alien, err error) {
	img, err := sdl.LoadBMP("sprites/aliencraft.bmp")
	if err != nil {
		return alien{}, fmt.Errorf("Init alien: %v", err)
	}
	defer img.Free()

	al.texture, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return alien{}, fmt.Errorf("Init player texutre: %v", err)
	}
	al.x = x
	al.y = y

	return al, nil
}

func (al *alien) draw(renderer *sdl.Renderer) {

	x := al.x - alienSize/2
	y := al.y - alienSize/2

	renderer.Copy(al.texture,
		&sdl.Rect{X: 0, Y: 0, W: 90, H: 90},
		&sdl.Rect{X: int32(x), Y: int32(y), W: 90, H: 90})
}
