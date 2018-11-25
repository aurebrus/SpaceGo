package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerTorpedoSize = 30
	torpedoSpeed      = 4
)

type playerTorpedo struct {
	texture *sdl.Texture
	x, y    float64
	active  bool
	angle   float64
}

func newPlayerTorpedo(renderer *sdl.Renderer) (tor playerTorpedo) {
	tor.texture = textureFromBMP(renderer, "sprites/playertorpedo.bmp")
	return tor
}

func (tor *playerTorpedo) draw(renderer *sdl.Renderer) {
	if !tor.active {
		return
	}
	x := tor.x - playerTorpedoSize/2
	y := tor.y - playerTorpedoSize/2

	renderer.Copy(tor.texture,
		&sdl.Rect{X: 0, Y: 0, W: playerTorpedoSize, H: playerTorpedoSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: playerTorpedoSize, H: playerTorpedoSize})
}

func (tor *playerTorpedo) update() {
	tor.x += torpedoSpeed * math.Cos(tor.angle)
	tor.y += torpedoSpeed * math.Sin(tor.angle)
}

var torpedoPool []*playerTorpedo

func initTorpedoPool(renderer *sdl.Renderer) {
	for i := 0; i < 40; i++ {
		tor := newPlayerTorpedo(renderer)
		torpedoPool = append(torpedoPool, &tor)
	}
}

func torpedoFromPool() (*playerTorpedo, bool) {
	for _, tor := range torpedoPool {
		if !tor.active {
			return tor, true
		}
	}
	return nil, false
}
