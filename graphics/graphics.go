package graphics

import (
	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

type Graphics struct {
	font     *ttf.Font
	Renderer *sdl.Renderer
}

func New(renderer *sdl.Renderer) *Graphics {
	ttf.Init()
	font, err := ttf.OpenFont("resources/Inconsolata-Regular.ttf", 24)
	if err != nil {
		panic("Unable to open font")
	}
	return &Graphics{
		font:     font,
		Renderer: renderer,
	}
}

func (g *Graphics) Print(s string) {
	surface := g.font.RenderText_Solid(s, sdl.Color{255, 255, 255, 255})
	texture := g.Renderer.CreateTextureFromSurface(surface)
	src := sdl.Rect{0, 0, surface.W, surface.H}
	dst := sdl.Rect{10, 10, surface.W, surface.H}
	g.Renderer.Copy(texture, &src, &dst)
}
