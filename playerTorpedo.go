package main

import "github.com/veandco/go-sdl2/sdl"

const (
	playerTorpedoSpeed = 1.0
	playerTorpedoSize  = 30
)

type playerTorpedo struct {
	texture *sdl.Texture
	x, y    float64
	active  bool
}

func newPlayerTorpedo(renderer *sdl.Renderer) (tor playerTorpedo) {
	tor.texture = textureFromBMP(renderer, "sprites/playertorpedo.bmp")
	return tor
}

func (tor *playerTorpedo) draw(renderer sdl.Renderer) {
	if !tor.active {
		return
	}
	x := tor.x - playerTorpedoSize/2
	y := tor.y - playerTorpedoSize/2

	renderer.Copy(tor.texture,
		&sdl.Rect{X: 0, Y: 0, W: playerTorpedoSize, H: playerTorpedoSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: playerTorpedoSize, H: playerTorpedoSize})

}

var torpedoPool []*playerTorpedo

func initTorpedoPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
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
