package main

import (
	"github.com/veandco/go-sdl2/sdl"

	"fps"
	"game"
	"graphics"
)

func main() {
	window := sdl.CreateWindow("gogame", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	renderer := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	renderer.SetDrawColor(50, 50, 50, 255)

	g := graphics.New(renderer)
	fps.Init(60, g)

	world := game.LoadWorld("resources/worlds/basic.txt", g)

	sword := game.NewSword(g)
	world.Player.SetActiveItem(sword)

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
					world.Player.Move(game.North)
				case sdl.K_d:
					world.Player.Move(game.East)
				case sdl.K_s:
					world.Player.Move(game.South)
				case sdl.K_a:
					world.Player.Move(game.West)
				case sdl.K_SPACE:
					world.Player.SetActiveItemState(true)
				}
			case *sdl.KeyUpEvent:
				switch event.Keysym.Sym {
				case sdl.K_w:
					world.Player.Stop(game.North)
				case sdl.K_d:
					world.Player.Stop(game.East)
				case sdl.K_s:
					world.Player.Stop(game.South)
				case sdl.K_a:
					world.Player.Stop(game.West)
				case sdl.K_SPACE:
					world.Player.SetActiveItemState(false)
				}
			}

		}

		dt := fps.Dt()
		world.Update(dt)

		g.Renderer.Clear()
		world.Draw()
		fps.DisplayFPS()
		g.Renderer.Present()

		fps.Update()
	}
	g.Renderer.Destroy()
	window.Destroy()
}
