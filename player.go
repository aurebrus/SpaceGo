package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed = 0.8
	playerSize  = 110
)

type player struct {
	texture *sdl.Texture
	x, y    float64
}

func newPlayer(renderer *sdl.Renderer) (plr player) {
	plr.texture = textureFromBMP(renderer, "sprites/playercraft.bmp")

	plr.x = float64(winW) / 2
	plr.y = float64(winH) - playerSize
	return plr
}

func (plr *player) draw(renderer *sdl.Renderer) {
	x := plr.x - playerSize/2
	y := plr.y - playerSize/2

	renderer.Copy(plr.texture,
		&sdl.Rect{X: 0, Y: 0, W: playerSize, H: playerSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: playerSize, H: playerSize})
}

func (plr *player) update() {
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_LEFT] == 1 {
		if plr.x-playerSize > 0 {
			plr.x -= playerSpeed
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if plr.x < winW-playerSize {
			plr.x += playerSpeed
		}
	}
}
