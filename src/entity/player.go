package entity

import (
	"github.com/veandco/go-sdl2/sdl"

	"graphics"
	"sprite"
	"world"
)

type Player struct {
	direction world.Direction
	speed     float32 // pixels/s
	spr       *sprite.Sprite
}

func NewPlayer(g *graphics.Graphics, w *world.World) *Player {
	player := &Player{
		direction: 0,
		spr:       sprite.New("resources/link.gif", g),
		speed:     200,
	}
	if startTile := w.FindTileKind(world.PlayerStart); startTile != nil {
		tRect := startTile.Bounds()
		player.spr.X = float32(tRect.X)
		player.spr.Y = float32(tRect.Y)
	}
	return player
}

func (p *Player) Move(d world.Direction) {
	p.direction |= d
}

func (p *Player) Stop(d world.Direction) {
	p.direction &^= d
}

func (p *Player) Update(dt uint32, w *world.World) {
	velocity := p.speed * float32(dt) / 1000
	if p.direction&world.North > 0 {
		p.spr.Y -= velocity
		w.CollideWithTiles(p, world.North)
	}
	if p.direction&world.East > 0 {
		p.spr.X += velocity
		w.CollideWithTiles(p, world.East)
	}
	if p.direction&world.South > 0 {
		p.spr.Y += velocity
		w.CollideWithTiles(p, world.South)
	}
	if p.direction&world.West > 0 {
		p.spr.X -= velocity
		w.CollideWithTiles(p, world.West)
	}
}

func (p *Player) Bounds() *sdl.Rect {
	return &sdl.Rect{int32(p.spr.X), int32(p.spr.Y), int32(p.spr.W), int32(p.spr.H)}
}

func (p *Player) SetBounds(r *sdl.Rect) {
	p.spr.X = float32(r.X)
	p.spr.Y = float32(r.Y)
}

func (p *Player) Draw() {
	p.spr.Draw()
}
