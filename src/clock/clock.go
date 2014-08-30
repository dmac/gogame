package clock

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"

	"graphics"
)

var g *graphics.Graphics
var maxFPS uint32       // e.g., 60fps
var startTick uint32    // time at which the current frame started
var lastSecTick uint32  // time of the last whole second
var frameCount uint32   // number of frames that have passed since lastSecTick
var frameDisplay uint32 // number of frames to display as current FPS
var dt uint32           // time at which Dt was last called

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

// Update updates the currently computed FPS and locks to framerate to maxFPS
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

// DisplayFPS prints the current FPS to the screen
func DisplayFPS() {
	g.Print(fmt.Sprintf("FPS:%d", frameDisplay))
}
