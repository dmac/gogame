package main

import (
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"

	"clock"
	"game"
	"graphics"
)

func main() {
	rand.Seed(time.Now().Unix())

	window := sdl.CreateWindow("gogame", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	renderer := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	renderer.SetDrawColor(50, 50, 50, 255)

	g := graphics.New(renderer, "resources/Inconsolata-Regular.ttf", 24, "resources/spritesheet.png", 32)
	clock.Init(60, g)

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

		dt := clock.Dt()
		world.Update(dt)

		g.Renderer.Clear()
		world.Draw()
		clock.DisplayFPS()
		g.Renderer.Present()

		clock.Update()
	}
	g.Renderer.Destroy()
	window.Destroy()
}
