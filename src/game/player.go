package game

import (
	"github.com/veandco/go-sdl2/sdl"

	"graphics"
	"sprite"
)

type Item interface {
	Bounded

	Activate()
	Deactivate()
	Draw()
}

type Player struct {
	x         float32
	y         float32
	speed     float32 // pixels/s
	direction Direction

	activeItem Item

	spr *sprite.Sprite
}

func NewPlayer(g *graphics.Graphics, w *World) *Player {
	player := &Player{
		speed:     200,
		direction: 0,
		spr:       sprite.New("resources/link.gif", g),
	}
	if startTile := w.FindTileKind(PlayerStart); startTile != nil {
		tRect := startTile.Bounds()
		player.x = float32(tRect.X)
		player.y = float32(tRect.Y)
	}
	return player
}

func (p *Player) Move(d Direction) {
	p.direction |= d
}

func (p *Player) Stop(d Direction) {
	p.direction &^= d
}

func (p *Player) Update(dt uint32, w *World) {
	velocity := p.speed * float32(dt) / 1000
	if p.direction&North > 0 {
		p.y -= velocity
		w.CollideWithTiles(p, North)
	}
	if p.direction&East > 0 {
		p.x += velocity
		w.CollideWithTiles(p, East)
	}
	if p.direction&South > 0 {
		p.y += velocity
		w.CollideWithTiles(p, South)
	}
	if p.direction&West > 0 {
		p.x -= velocity
		w.CollideWithTiles(p, West)
	}
}

func (p *Player) SetActiveItem(item Item) {
	p.activeItem = item
}

func (p *Player) SetActiveItemState(s bool) {
	if s {
		p.activeItem.Activate()
	} else {
		p.activeItem.Deactivate()
	}
}

func (p *Player) Bounds() *sdl.Rect {
	return &sdl.Rect{int32(p.x), int32(p.y), int32(p.spr.W), int32(p.spr.H)}
}

func (p *Player) SetBounds(r *sdl.Rect) {
	p.x = float32(r.X)
	p.y = float32(r.Y)
}

func (p *Player) Draw() {
	p.spr.X = p.x
	p.spr.Y = p.y
	p.spr.Draw()

	p.activeItem.SetBounds(&sdl.Rect{int32(p.x) + 15, int32(p.y) + 30, 0, 0})
	p.activeItem.Draw()
}
