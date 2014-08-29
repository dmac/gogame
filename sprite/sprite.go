package sprite

import (
	"github.com/veandco/go-sdl2/sdl"
	img "github.com/veandco/go-sdl2/sdl_image"

	"../graphics"
)

type Sprite struct {
	X       int32
	Y       int32
	W       int32
	H       int32
	texture *sdl.Texture
	g       *graphics.Graphics
}

func New(filename string, graphics *graphics.Graphics) *Sprite {
	surface := img.Load("resources/link.gif")
	texture := graphics.Renderer.CreateTextureFromSurface(surface)
	return &Sprite{
		X:       100,
		Y:       100,
		W:       surface.W,
		H:       surface.H,
		texture: texture,
		g:       graphics,
	}
}

func (s *Sprite) Draw() {
	src := sdl.Rect{0, 0, s.W, s.H}
	dst := sdl.Rect{s.X, s.Y, s.W, s.H}
	s.g.Renderer.Copy(s.texture, &src, &dst)
}
