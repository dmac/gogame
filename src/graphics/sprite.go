package graphics

import (
	"github.com/veandco/go-sdl2/sdl"
	img "github.com/veandco/go-sdl2/sdl_image"
)

type spritesheet struct {
	size uint
	g    *Graphics
	t    *sdl.Texture
}

func newSpritesheet(filename string, size uint, g *Graphics) *spritesheet {
	surface := img.Load(filename)
	texture := g.Renderer.CreateTextureFromSurface(surface)
	return &spritesheet{
		size: size,
		g:    g,
		t:    texture,
	}
}

type Sprite struct {
	X   float32 // world coords
	Y   float32 // world coords
	W   uint
	H   uint
	ssx uint // sprite sheet coords
	ssy uint // sprite sheet coords
	ss  *spritesheet
}

func NewSprite(row, col, nrows, ncols uint, g *Graphics) *Sprite {
	return &Sprite{
		X:   0,
		Y:   0,
		W:   ncols * g.ss.size,
		H:   nrows * g.ss.size,
		ssx: col * g.ss.size,
		ssy: row * g.ss.size,
		ss:  g.ss,
	}
}

func (s *Sprite) Draw() {
	src := sdl.Rect{int32(s.ssx), int32(s.ssy), int32(s.W), int32(s.H)}
	dst := sdl.Rect{int32(s.X), int32(s.Y), int32(s.W), int32(s.H)}
	s.ss.g.Renderer.Copy(s.ss.t, &src, &dst)
}
