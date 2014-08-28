package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	img "github.com/veandco/go-sdl2/sdl_image"
)

type Sprite struct {
	x       int32
	y       int32
	w       int32
	h       int32
	texture *sdl.Texture
}

func NewSprite(renderer *sdl.Renderer, filename string) *Sprite {
	surface := img.Load("resources/link.gif")
	texture := renderer.CreateTextureFromSurface(surface)
	return &Sprite{
		x:       100,
		y:       100,
		w:       surface.W,
		h:       surface.H,
		texture: texture,
	}
}

func (s *Sprite) Draw(renderer *sdl.Renderer) {
	src := sdl.Rect{0, 0, s.w, s.h}
	dst := sdl.Rect{s.x, s.y, s.w, s.h}
	renderer.Copy(s.texture, &src, &dst)
}

func main() {
	window := sdl.CreateWindow("gogame", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	renderer := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	sprite := NewSprite(renderer, "resources/link.gif")

	moveUp := false
	moveRight := false
	moveDown := false
	moveLeft := false

	running := true
	for running {
		renderer.Clear()
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

		if moveUp {
			sprite.y -= 1
		}
		if moveRight {
			sprite.x += 1
		}
		if moveDown {
			sprite.y += 1
		}
		if moveLeft {
			sprite.x -= 1
		}

		sprite.Draw(renderer)
		renderer.Present()
	}
	renderer.Destroy()
	window.Destroy()
}
