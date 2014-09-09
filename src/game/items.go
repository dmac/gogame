package game

import (
	"github.com/veandco/go-sdl2/sdl"

	"graphics"
)

type Sword struct {
	x       float32
	y       float32
	faceDir Direction
	active  bool
	damage  int32
	sprs    []*graphics.Sprite
}

func NewSword(g *graphics.Graphics) *Sword {
	return &Sword{
		damage: 10,
		sprs: []*graphics.Sprite{
			graphics.NewSprite(2, 2, 1, 1, g),
			graphics.NewSprite(2, 3, 1, 1, g),
			graphics.NewSprite(2, 0, 1, 1, g),
			graphics.NewSprite(2, 1, 1, 1, g),
		},
	}
}

func (s *Sword) Activate() {
	s.active = true
}

func (s *Sword) Deactivate() {
	s.active = false
}

func (s *Sword) Update(dt uint32, w *World) {
	for i := range w.Enemies {
		sRect := s.Bounds()
		eRect := w.Enemies[i].Bounds()
		if s.active && sRect.HasIntersection(eRect) {
			w.Enemies[i].ChangeHealth(-1 * s.damage)
		}
	}
}

func (s *Sword) Draw() {
	if !s.active {
		return
	}
	var spr *graphics.Sprite
	switch s.faceDir {
	case North:
		spr = s.sprs[0]
	case East:
		spr = s.sprs[1]
	case South:
		spr = s.sprs[2]
	case West:
		spr = s.sprs[3]
	}
	spr.X = s.x
	spr.Y = s.y
	spr.Draw()
}

func (s *Sword) SetFaceDir(d Direction) {
	s.faceDir = d
}

func (s *Sword) Bounds() *sdl.Rect {
	return &sdl.Rect{int32(s.x), int32(s.y), int32(s.sprs[0].W), int32(s.sprs[0].H)}
}

func (s *Sword) SetBounds(b *sdl.Rect) {
	s.x = float32(b.X)
	s.y = float32(b.Y)
}
