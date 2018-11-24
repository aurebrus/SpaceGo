package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	alienTorpedoSize = 30
	torpedoASpeed    = 4
)

type alienTorpedo struct {
	texture *sdl.Texture
	x, y    float64
	active  bool
	angle   float64
}

func newAlienTorpedo(renderer *sdl.Renderer) (tor alienTorpedo) {
	tor.texture = textureFromBMP(renderer, "sprites/alientorpedo.bmp")
	return tor
}

func (tor *alienTorpedo) draw(renderer *sdl.Renderer) {
	if !tor.active {
		return
	}
	x := tor.x - playerTorpedoSize/2
	y := tor.y - playerTorpedoSize/2

	renderer.Copy(tor.texture,
		&sdl.Rect{X: 0, Y: 0, W: alienTorpedoSize, H: alienTorpedoSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: alienTorpedoSize, H: alienTorpedoSize})
}

func (tor *alienTorpedo) update() {
	tor.x -= torpedoASpeed * math.Cos(tor.angle)
	tor.y -= torpedoASpeed * math.Sin(tor.angle)
}

var torpedoAPool []*alienTorpedo

func initATorpedoPool(renderer *sdl.Renderer) {
	for i := 0; i < 300; i++ {
		tor := newAlienTorpedo(renderer)
		torpedoAPool = append(torpedoAPool, &tor)
	}
}

func torpedoAFromPool() (*alienTorpedo, bool) {
	for _, tor := range torpedoAPool {
		if !tor.active {
			return tor, true
		}
	}
	return nil, false
}
