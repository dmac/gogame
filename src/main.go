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
	player := game.NewPlayer(g, world)
	moblin := game.NewMoblin(g, world)
	sword := game.NewSword(g)

	player.SetActiveItem(sword)

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
					player.Move(game.North)
				case sdl.K_d:
					player.Move(game.East)
				case sdl.K_s:
					player.Move(game.South)
				case sdl.K_a:
					player.Move(game.West)
				case sdl.K_SPACE:
					player.SetActiveItemState(true)
				}
			case *sdl.KeyUpEvent:
				switch event.Keysym.Sym {
				case sdl.K_w:
					player.Stop(game.North)
				case sdl.K_d:
					player.Stop(game.East)
				case sdl.K_s:
					player.Stop(game.South)
				case sdl.K_a:
					player.Stop(game.West)
				case sdl.K_SPACE:
					player.SetActiveItemState(false)
				}
			}

		}

		dt := fps.Dt()
		player.Update(dt, world)
		moblin.Update(dt, world)

		g.Renderer.Clear()
		player.Draw()
		moblin.Draw()
		world.Draw()
		fps.DisplayFPS()
		g.Renderer.Present()

		fps.Update()
	}
	g.Renderer.Destroy()
	window.Destroy()
}
