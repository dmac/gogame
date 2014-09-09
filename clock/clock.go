package clock

import "github.com/veandco/go-sdl2/sdl"

var max uint32         // e.g., 60fps
var fps uint32         // number of frames to display as current FPS
var startTick uint32   // time at which the current frame started
var lastSecTick uint32 // time of the last whole second
var frameCount uint32  // number of frames that have passed since lastSecTick
var dt uint32          // time at which Dt was last called

func Init(maxFPS uint32) {
	tick := sdl.GetTicks()
	max = maxFPS
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

// Update updates the currently computed FPS and locks to framerate to max
func Update() {
	frameCount += 1
	if float64(startTick-lastSecTick)/1000 > 1 {
		lastSecTick = startTick
		fps = frameCount
		frameCount = 0
	}

	delay := float32(1000/max) - float32(sdl.GetTicks()-startTick)
	if delay > 0 {
		sdl.Delay(uint32(delay))
	}
	startTick = sdl.GetTicks()
}

// FPS returns the currently calculated frames per second
func FPS() uint32 {
	return fps
}
