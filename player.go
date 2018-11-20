package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed = 0.35
	playerSize  = 120
)

type player struct {
	texture *sdl.Texture
	x, y    float64
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

	plr.x = float64(winW) / 2
	plr.y = float64(winH) - 150.0
	return plr, nil
}

func (plr *player) draw(renderer *sdl.Renderer) {

	renderer.Copy(plr.texture,
		&sdl.Rect{X: 0, Y: 0, W: 115, H: 115},
		&sdl.Rect{X: int32(plr.x), Y: int32(plr.y), W: 110, H: 110})
}

func (plr *player) update() {
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_LEFT] == 1 {
		plr.x -= playerSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		plr.x += playerSpeed
	}
}
