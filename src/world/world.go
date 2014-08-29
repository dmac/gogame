package world

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/veandco/go-sdl2/sdl"

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

type Bounded interface {
	Bounds() *sdl.Rect
	SetBounds(*sdl.Rect)
}

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

func LoadWorld(filename string, g *graphics.Graphics) *World {
	f, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Unable to open world file: %s", filename))
	}
	defer f.Close()

	tiles := make([]tile, 0)

	r := bufio.NewReader(f)
	line, isPrefix, err := r.ReadLine()
	row := 0
	for err != io.EOF {
		if isPrefix {
			panic("World loader buffer size too small")
		}

		s := string(line)
		for col, c := range s {
			switch c {
			case '-', '|':
				spr := sprite.New("resources/block.gif", g)
				spr.X = float32(col) * spr.W
				spr.Y = float32(row) * spr.H
				newTile := tile{
					row:  int32(row),
					col:  int32(col),
					kind: Wall,
					spr:  spr,
				}
				tiles = append(tiles, newTile)
			}
		}

		row += 1
		line, isPrefix, err = r.ReadLine()
	}

	return &World{
		tiles: tiles,
	}
}

func (w *World) CollideWithTiles(b Bounded, d Direction) {
	for _, t := range w.tiles {
		bRect := b.Bounds()
		tRect := t.Bounds()
		if bRect.HasIntersection(tRect) {
			switch d {
			case North:
				for bRect.Y < tRect.Y+tRect.H {
					bRect.Y += 1
				}
			case East:
				for bRect.X+bRect.W > tRect.X {
					bRect.X -= 1
				}
			case South:
				for bRect.Y+bRect.H > tRect.Y {
					bRect.Y -= 1
				}
			case West:
				for bRect.X < tRect.X+tRect.W {
					bRect.X += 1
				}
			}
			b.SetBounds(bRect)
		}
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

func (t *tile) Bounds() *sdl.Rect {
	if t.spr == nil {
		return nil
	}
	return &sdl.Rect{int32(t.spr.X), int32(t.spr.Y), int32(t.spr.W), int32(t.spr.H)}
}
