package game

import (
	"github.com/veandco/go-sdl2/sdl"

	"graphics"
)

type Item interface {
	Bounded

	Activate()
	Deactivate()
	Update(dt uint32, w *World)
	Draw()
}

type Player struct {
	x         float32
	y         float32
	speed     float32 // pixels/s
	direction Direction

	activeItem Item

	lastSpr *graphics.Sprite
	sprs    []*graphics.Sprite
}

func NewPlayer(g *graphics.Graphics) *Player {
	player := &Player{
		speed:     200,
		direction: 0,
		sprs: []*graphics.Sprite{
			graphics.NewSprite(0, 2, 1, 1, g),
			graphics.NewSprite(0, 3, 1, 1, g),
			graphics.NewSprite(0, 0, 1, 1, g),
			graphics.NewSprite(0, 1, 1, 1, g),
		},
	}
	player.lastSpr = player.sprs[2]
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
	p.activeItem.Update(dt, w)
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
	return &sdl.Rect{int32(p.x), int32(p.y), int32(p.sprs[0].W), int32(p.sprs[0].H)}
}

func (p *Player) SetBounds(r *sdl.Rect) {
	p.x = float32(r.X)
	p.y = float32(r.Y)
}

func (p *Player) Draw() {
	spr := p.lastSpr
	if p.direction&North > 0 {
		spr = p.sprs[0]
	}
	if p.direction&East > 0 {
		spr = p.sprs[1]
	}
	if p.direction&South > 0 {
		spr = p.sprs[2]
	}
	if p.direction&West > 0 {
		spr = p.sprs[3]
	}
	p.lastSpr = spr
	spr.X = p.x
	spr.Y = p.y
	spr.Draw()

	p.activeItem.SetBounds(&sdl.Rect{int32(p.x) + 9, int32(p.y) + 25, 0, 0})
	p.activeItem.Draw()
}
