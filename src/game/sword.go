package game

import (
	"github.com/veandco/go-sdl2/sdl"

	"graphics"
	"sprite"
)

type Sword struct {
	x      float32
	y      float32
	active bool
	damage int32
	spr    *sprite.Sprite
}

func NewSword(g *graphics.Graphics) *Sword {
	return &Sword{
		damage: 10,
		spr:    sprite.New("resources/sword.gif", g),
	}
}

func (s *Sword) Activate() {
	s.active = true
}

func (s *Sword) Deactivate() {
	s.active = false
}

func (s *Sword) Update(dt uint32, w *World) {
	if w.Moblin == nil {
		return
	}
	sRect := s.Bounds()
	mRect := w.Moblin.Bounds()
	if s.active && sRect.HasIntersection(mRect) {
		w.Moblin.ChangeHealth(-1 * s.damage)
	}
}

func (s *Sword) Draw() {
	if !s.active {
		return
	}
	s.spr.X = s.x
	s.spr.Y = s.y
	s.spr.Draw()
}

func (s *Sword) Bounds() *sdl.Rect {
	return &sdl.Rect{int32(s.x), int32(s.y), int32(s.spr.W), int32(s.spr.H)}
}

func (s *Sword) SetBounds(b *sdl.Rect) {
	s.x = float32(b.X)
	s.y = float32(b.Y)
}