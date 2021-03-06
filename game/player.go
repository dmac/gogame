package game

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/dmac/gogame/graphics"
)

type Item interface {
	Bounded

	Activate()
	Deactivate()
	Update(dt uint32, w *World)
	SetFaceDir(d Direction)
	Draw()
}

type Player struct {
	x          float32
	y          float32
	speed      float32 // pixels/s
	moveDir    Direction
	faceDir    Direction
	activeItem Item
	sprs       []*graphics.Sprite
}

func NewPlayer(g *graphics.Graphics) *Player {
	player := &Player{
		speed:   200,
		faceDir: South,
		sprs: []*graphics.Sprite{
			graphics.NewSprite(0, 2, 1, 1, g),
			graphics.NewSprite(0, 3, 1, 1, g),
			graphics.NewSprite(0, 0, 1, 1, g),
			graphics.NewSprite(0, 1, 1, 1, g),
		},
	}
	return player
}

func (p *Player) Move(d Direction) {
	p.moveDir |= d
}

func (p *Player) Stop(d Direction) {
	p.moveDir &^= d
}

func (p *Player) Update(dt uint32, w *World) {
	velocity := p.speed * float32(dt) / 1000
	if p.moveDir&North > 0 {
		p.y -= velocity
		w.CollideWithTiles(p, North)
	}
	if p.moveDir&East > 0 {
		p.x += velocity
		w.CollideWithTiles(p, East)
	}
	if p.moveDir&South > 0 {
		p.y += velocity
		w.CollideWithTiles(p, South)
	}
	if p.moveDir&West > 0 {
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
	var spr *graphics.Sprite
	var itemDx int32 = 0
	var itemDy int32 = 0

	if p.moveDir == North || p.moveDir == East || p.moveDir == South || p.moveDir == West {
		p.faceDir = p.moveDir
	}

	switch p.faceDir {
	case North:
		spr = p.sprs[0]
		itemDx = -8
		itemDy = -9
	case East:
		spr = p.sprs[1]
		itemDx = 13
		itemDy = 9
	case South:
		spr = p.sprs[2]
		itemDx = 9
		itemDy = 25
	case West:
		spr = p.sprs[3]
		itemDx = -13
		itemDy = 8
	}

	spr.X = p.x
	spr.Y = p.y
	spr.Draw()

	p.activeItem.SetFaceDir(p.faceDir)
	p.activeItem.SetBounds(&sdl.Rect{int32(p.x) + itemDx, int32(p.y) + itemDy, 0, 0})
	p.activeItem.Draw()
}
