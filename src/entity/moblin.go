package entity

import (
	"github.com/veandco/go-sdl2/sdl"

	"graphics"
	"sprite"
	"world"
)

type Moblin struct {
	direction world.Direction
	speed     float32 // pixels/s
	spr       *sprite.Sprite
}

func NewMoblin(g *graphics.Graphics, w *world.World) *Moblin {
	moblin := &Moblin{
		direction: world.South,
		spr:       sprite.New("resources/moblin.gif", g),
		speed:     50,
	}
	if startTile := w.FindTileKind(world.MoblinStart); startTile != nil {
		tRect := startTile.Bounds()
		moblin.spr.X = float32(tRect.X)
		moblin.spr.Y = float32(tRect.Y)
	}
	return moblin
}

func (m *Moblin) Update(dt uint32, w *world.World) {
	velocity := m.speed * float32(dt) / 1000
	if m.direction&world.North > 0 {
		m.spr.Y -= velocity
		w.CollideWithTiles(m, world.North)
	}
	if m.direction&world.East > 0 {
		m.spr.X += velocity
		w.CollideWithTiles(m, world.East)
	}
	if m.direction&world.South > 0 {
		m.spr.Y += velocity
		w.CollideWithTiles(m, world.South)
	}
	if m.direction&world.West > 0 {
		m.spr.X -= velocity
		w.CollideWithTiles(m, world.West)
	}
}

func (m *Moblin) Bounds() *sdl.Rect {
	return &sdl.Rect{int32(m.spr.X), int32(m.spr.Y), int32(m.spr.W), int32(m.spr.H)}
}

func (m *Moblin) SetBounds(r *sdl.Rect) {
	m.spr.X = float32(r.X)
	m.spr.Y = float32(r.Y)
}

func (m *Moblin) Draw() {
	m.spr.Draw()
}
