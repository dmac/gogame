package graphics

import (
	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

type Graphics struct {
	Renderer *sdl.Renderer
	font     *ttf.Font
	ss       *spritesheet
}

func New(renderer *sdl.Renderer, fontFilename string, fontSize int,
	spritesheetFilename string, spritesheetSize uint) *Graphics {

	ttf.Init()
	font, err := ttf.OpenFont(fontFilename, fontSize)
	if err != nil {
		panic("Unable to open font")
	}
	g := &Graphics{
		Renderer: renderer,
		font:     font,
	}
	spritesheet := newSpritesheet(spritesheetFilename, spritesheetSize, g)
	g.ss = spritesheet
	return g
}

func (g *Graphics) Print(s string) {
	surface := g.font.RenderText_Solid(s, sdl.Color{255, 255, 255, 255})
	texture := g.Renderer.CreateTextureFromSurface(surface)
	src := sdl.Rect{0, 0, surface.W, surface.H}
	dst := sdl.Rect{10, 10, surface.W, surface.H}
	g.Renderer.Copy(texture, &src, &dst)
}
