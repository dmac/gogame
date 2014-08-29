package sprite

import (
	"github.com/veandco/go-sdl2/sdl"
	img "github.com/veandco/go-sdl2/sdl_image"

	"graphics"
)

type Sprite struct {
	X       float32
	Y       float32
	W       float32
	H       float32
	texture *sdl.Texture
	g       *graphics.Graphics
}

func New(filename string, graphics *graphics.Graphics) *Sprite {
	surface := img.Load(filename)
	texture := graphics.Renderer.CreateTextureFromSurface(surface)
	return &Sprite{
		X:       0,
		Y:       0,
		W:       float32(surface.W),
		H:       float32(surface.H),
		texture: texture,
		g:       graphics,
	}
}

func (s *Sprite) Draw() {
	src := sdl.Rect{0, 0, int32(s.W), int32(s.H)}
	dst := sdl.Rect{int32(s.X), int32(s.Y), int32(s.W), int32(s.H)}
	s.g.Renderer.Copy(s.texture, &src, &dst)
}
