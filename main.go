package main

import (
	"github.com/veandco/go-sdl2/sdl"

	"./entity"
	"./fps"
	"./graphics"
)

func main() {
	window := sdl.CreateWindow("gogame", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	renderer := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	renderer.SetDrawColor(50, 50, 50, 255)

	g := graphics.New(renderer)
	fps.Init(60, g)

	player := entity.NewPlayer(g)

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
					player.Move(entity.North)
				case sdl.K_d:
					player.Move(entity.East)
				case sdl.K_s:
					player.Move(entity.South)
				case sdl.K_a:
					player.Move(entity.West)
				}
			case *sdl.KeyUpEvent:
				switch event.Keysym.Sym {
				case sdl.K_w:
					player.Stop(entity.North)
				case sdl.K_d:
					player.Stop(entity.East)
				case sdl.K_s:
					player.Stop(entity.South)
				case sdl.K_a:
					player.Stop(entity.West)
				}
			}

		}

		dt := fps.Dt()
		player.Update(dt)

		g.Renderer.Clear()

		player.Draw()
		fps.DisplayFPS()

		g.Renderer.Present()

		fps.Update()
	}
	renderer.Destroy()
	window.Destroy()
}
