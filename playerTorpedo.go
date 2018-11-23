package main

import "github.com/veandco/go-sdl2/sdl"

const (
	playerTorpedoSpeed = 1.0
	playerTorpedoSize  = 30
)

type playerTorpedo struct {
	texture *sdl.Texture
	x, y    float64
}

func newPlayerTorpedo(renderer *sdl.Renderer) (tor playerTorpedo) {
	tor.texture = textureFromBMP(renderer, "sprites/playertorpedo.bmp")
	return tor
}

func (tor *playerTorpedo) draw(renderer sdl.Renderer) {
	x := tor.x - playerTorpedoSize/2
	y := tor.y - playerTorpedoSize/2

	renderer.Copy(tor.texture,
		&sdl.Rect{X: 0, Y: 0, W: playerTorpedoSize, H: playerTorpedoSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: playerTorpedoSize, H: playerTorpedoSize})

}
