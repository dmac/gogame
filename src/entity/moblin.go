package entity

import (
	"github.com/veandco/go-sdl2/sdl"

	"graphics"
	"sprite"
	"world"
)

type Moblin struct {
	x         float32
	y         float32
	speed     float32 // pixels/s
	direction world.Direction
	spr       *sprite.Sprite
}

func NewMoblin(g *graphics.Graphics, w *world.World) *Moblin {
	moblin := &Moblin{
		speed:     50,
		direction: world.South,
		spr:       sprite.New("resources/moblin.gif", g),
	}
	if startTile := w.FindTileKind(world.MoblinStart); startTile != nil {
		tRect := startTile.Bounds()
		moblin.x = float32(tRect.X)
		moblin.y = float32(tRect.Y)
	}
	return moblin
}

func (m *Moblin) Update(dt uint32, w *world.World) {
	velocity := m.speed * float32(dt) / 1000
	if m.direction&world.North > 0 {
		m.y -= velocity
		w.CollideWithTiles(m, world.North)
	}
	if m.direction&world.East > 0 {
		m.x += velocity
		w.CollideWithTiles(m, world.East)
	}
	if m.direction&world.South > 0 {
		m.y += velocity
		w.CollideWithTiles(m, world.South)
	}
	if m.direction&world.West > 0 {
		m.x -= velocity
		w.CollideWithTiles(m, world.West)
	}
}

func (m *Moblin) Bounds() *sdl.Rect {
	return &sdl.Rect{int32(m.x), int32(m.y), int32(m.spr.W), int32(m.spr.H)}
}

func (m *Moblin) SetBounds(r *sdl.Rect) {
	m.x = float32(r.X)
	m.y = float32(r.Y)
}

func (m *Moblin) Draw() {
	m.spr.X = m.x
	m.spr.Y = m.y
	m.spr.Draw()
}
