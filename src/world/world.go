package world

import (
	"graphics"
	"sprite"
)

type tileKind struct{ v byte }

var Wall tileKind = tileKind{v: 0}
var PlayerStart tileKind = tileKind{v: 1}

type tile struct {
	row  int32
	col  int32
	kind tileKind
	spr  *sprite.Sprite
}

type World struct {
	tiles []tile
}

func NewWorld(g *graphics.Graphics) *World {
	wallSprite := sprite.New("resources/block.gif", g)
	return &World{
		tiles: []tile{
			tile{
				row:  5,
				col:  3,
				kind: Wall,
				spr:  wallSprite,
			},
			tile{
				row:  10,
				col:  10,
				kind: PlayerStart,
			},
		},
	}
}

func (w *World) Draw() {
	for _, tile := range w.tiles {
		tile.Draw()
	}
}

func (t *tile) Draw() {
	if t.spr == nil {
		return
	}
	t.spr.X = float32(t.col) * t.spr.W
	t.spr.Y = float32(t.row) * t.spr.H
	t.spr.Draw()
}
