package entity

import (
	"graphics"
	"sprite"
)

type Direction byte

const (
	North Direction = 1 << iota
	East
	South
	West
)

type Player struct {
	Direction Direction
	s         *sprite.Sprite
	speed     float32 // pixels/s
}

func NewPlayer(g *graphics.Graphics) *Player {
	return &Player{
		Direction: 0,
		s:         sprite.New("resources/link.gif", g),
		speed:     200,
	}
}

func (p *Player) Move(d Direction) {
	p.Direction |= d
}

func (p *Player) Stop(d Direction) {
	p.Direction &^= d
}

func (p *Player) Update(dt uint32) {
	velocity := p.speed * float32(dt) / 1000
	if p.Direction&North > 0 {
		p.s.Y -= velocity
	}
	if p.Direction&East > 0 {
		p.s.X += velocity
	}
	if p.Direction&South > 0 {
		p.s.Y += velocity
	}
	if p.Direction&West > 0 {
		p.s.X -= velocity
	}
}

func (p *Player) Draw() {
	p.s.Draw()
}