package entity

import (
	"../graphics"
	"../sprite"
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
	speed     uint
}

func NewPlayer(g *graphics.Graphics) *Player {
	return &Player{
		Direction: 0,
		s:         sprite.New("resources/link.gif", g),
		speed:     10,
	}
}

func (p *Player) Move(d Direction) {
	p.Direction |= d
}

func (p *Player) Stop(d Direction) {
	p.Direction &^= d
}

// TODO: Take dt and interpolate with speed
func (p *Player) Update() {
	if p.Direction&North > 0 {
		p.s.Y -= int32(p.speed)
	}
	if p.Direction&East > 0 {
		p.s.X += int32(p.speed)
	}
	if p.Direction&South > 0 {
		p.s.Y += int32(p.speed)
	}
	if p.Direction&West > 0 {
		p.s.X -= int32(p.speed)
	}
}

func (p *Player) Draw() {
	p.s.Draw()
}
