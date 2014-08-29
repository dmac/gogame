package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"

	"./fps"
	"./graphics"
	"./sprite"
)

func main() {
	window := sdl.CreateWindow("gogame", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	renderer := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	g := graphics.New(renderer)

	sprite := sprite.New("resources/link.gif", g)

	moveUp := false
	moveRight := false
	moveDown := false
	moveLeft := false

	fps.Init(60, g)

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
			sprite.Y -= speed
		}
		if moveRight {
			sprite.X += speed
		}
		if moveDown {
			sprite.Y += speed
		}
		if moveLeft {
			sprite.X -= speed
		}

		g.Renderer.Clear()

		sprite.Draw()
		fps.DisplayFPS()

		g.Renderer.Present()

		fps.Update()
	}
	renderer.Destroy()
	window.Destroy()
}
