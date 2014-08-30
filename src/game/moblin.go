package game

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"

	"graphics"
	"sprite"
)

type Moblin struct {
	x         float32
	y         float32
	speed     float32 // pixels/s
	health    int32
	maxHealth int32
	direction Direction
	spr       *sprite.Sprite
}

func NewMoblin(g *graphics.Graphics) *Moblin {
	return &Moblin{
		health:    100,
		maxHealth: 100,
		speed:     50,
		direction: South,
		spr:       sprite.New("resources/moblin.gif", g),
	}
}

func (m *Moblin) Update(dt uint32, w *World) {
	velocity := m.speed * float32(dt) / 1000
	if m.direction&North > 0 {
		m.y -= velocity
		w.CollideWithTiles(m, North)
	}
	if m.direction&East > 0 {
		m.x += velocity
		w.CollideWithTiles(m, East)
	}
	if m.direction&South > 0 {
		m.y += velocity
		w.CollideWithTiles(m, South)
	}
	if m.direction&West > 0 {
		m.x -= velocity
		w.CollideWithTiles(m, West)
	}
}

func (m *Moblin) ChangeHealth(amount int32) {
	m.health += amount
	fmt.Printf("Moblin Health: %d/%d\n", m.health, m.maxHealth)
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
