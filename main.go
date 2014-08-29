package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	img "github.com/veandco/go-sdl2/sdl_image"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

type FPS struct {
	text         *Text
	maxFPS       uint32
	startTick    uint32
	lastSecTick  uint32
	frameCount   uint32
	frameDisplay uint32
}

func newFPS(maxFPS uint32, text *Text) *FPS {
	tick := sdl.GetTicks()
	return &FPS{
		text:        text,
		maxFPS:      maxFPS,
		startTick:   tick,
		lastSecTick: tick,
	}
}

func (fps *FPS) update() {
	fps.frameCount += 1
	if float64(fps.startTick-fps.lastSecTick)/1000 > 1 {
		fps.lastSecTick = fps.startTick
		fps.frameDisplay = fps.frameCount
		fps.frameCount = 0
	}

	delay := float32(1000/fps.maxFPS) - float32(sdl.GetTicks()-fps.startTick)
	if delay > 0 {
		sdl.Delay(uint32(delay))
	}
	fps.startTick = sdl.GetTicks()
}

func (fps *FPS) displayFPS(r *sdl.Renderer) {
	fps.text.drawText(fmt.Sprintf("FPS:%d", fps.frameDisplay), r)
}

type Sprite struct {
	x       int32
	y       int32
	w       int32
	h       int32
	texture *sdl.Texture
}

func NewSprite(r *sdl.Renderer, filename string) *Sprite {
	surface := img.Load("resources/link.gif")
	texture := r.CreateTextureFromSurface(surface)
	return &Sprite{
		x:       100,
		y:       100,
		w:       surface.W,
		h:       surface.H,
		texture: texture,
	}
}

func (s *Sprite) draw(r *sdl.Renderer) {
	src := sdl.Rect{0, 0, s.w, s.h}
	dst := sdl.Rect{s.x, s.y, s.w, s.h}
	r.Copy(s.texture, &src, &dst)
}

type Text struct {
	font *ttf.Font
}

func (t *Text) drawText(s string, r *sdl.Renderer) {
	surface := t.font.RenderText_Solid(s, sdl.Color{255, 255, 255, 255})
	texture := r.CreateTextureFromSurface(surface)
	src := sdl.Rect{0, 0, surface.W, surface.H}
	dst := sdl.Rect{10, 10, surface.W, surface.H}
	r.Copy(texture, &src, &dst)
}

func main() {
	window := sdl.CreateWindow("gogame", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	renderer := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	ttf.Init()
	font, err := ttf.OpenFont("resources/Inconsolata-Regular.ttf", 24)
	if err != nil {
		panic("Unable to open font")
	}
	text := Text{font: font}

	sprite := NewSprite(renderer, "resources/link.gif")

	moveUp := false
	moveRight := false
	moveDown := false
	moveLeft := false

	fps := newFPS(60, &text)

	running := true
	for running {
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch event := e.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyDownEvent:
				switch event.Keysym.Sym {
				case sdl.K_ESCAPE:
					running = false
				case sdl.K_w:
					moveUp = true
				case sdl.K_d:
					moveRight = true
				case sdl.K_s:
					moveDown = true
				case sdl.K_a:
					moveLeft = true
				default:
					fmt.Println(event.Keysym)
				}
			case *sdl.KeyUpEvent:
				switch event.Keysym.Sym {
				case sdl.K_w:
					moveUp = false
				case sdl.K_d:
					moveRight = false
				case sdl.K_s:
					moveDown = false
				case sdl.K_a:
					moveLeft = false
				}
			}

		}

		var speed int32 = 10
		if moveUp {
			sprite.y -= speed
		}
		if moveRight {
			sprite.x += speed
		}
		if moveDown {
			sprite.y += speed
		}
		if moveLeft {
			sprite.x -= speed
		}

		renderer.Clear()

		sprite.draw(renderer)
		fps.displayFPS(renderer)

		renderer.Present()

		fps.update()
	}
	renderer.Destroy()
	window.Destroy()
}
