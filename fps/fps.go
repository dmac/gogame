package fps

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"

	"../graphics"
)

var g *graphics.Graphics
var maxFPS uint32
var startTick uint32
var lastSecTick uint32
var frameCount uint32
var frameDisplay uint32
var dt uint32

func Init(max uint32, graphics *graphics.Graphics) {
	tick := sdl.GetTicks()
	g = graphics
	maxFPS = max
	startTick = tick
	lastSecTick = tick
	dt = tick
}

// Dt returns the time between now and the last time it was called in milliseconds.
func Dt() uint32 {
	now := sdl.GetTicks()
	delta := now - dt
	dt = now
	return delta
}

func Update() {
	frameCount += 1
	if float64(startTick-lastSecTick)/1000 > 1 {
		lastSecTick = startTick
		frameDisplay = frameCount
		frameCount = 0
	}

	delay := float32(1000/maxFPS) - float32(sdl.GetTicks()-startTick)
	if delay > 0 {
		sdl.Delay(uint32(delay))
	}
	startTick = sdl.GetTicks()
}

func DisplayFPS() {
	g.Print(fmt.Sprintf("FPS:%d", frameDisplay))
}
