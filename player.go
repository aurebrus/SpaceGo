package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	texture *sdl.Texture
}

func newPlayer(renderer *sdl.Renderer) (plr player, err error) {

	img, err := sdl.LoadBMP("sprites/playercraft.bmp")
	if err != nil {
		return player{}, fmt.Errorf("Init player: %v", err)
	}
	defer img.Free()

	plr.texture, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return player{}, fmt.Errorf("Init player texutre: %v", err)
	}

	return plr, nil
}

func (plr *player) draw(renderer *sdl.Renderer) {
	renderer.Copy(plr.texture,
		&sdl.Rect{X: 0, Y: 0, W: 115, H: 115},
		&sdl.Rect{X: 0, Y: 0, W: 115, H: 115})
}
