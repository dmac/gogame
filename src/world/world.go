package world

import (
	"bufio"
	"fmt"
	"io"
	"os"

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

func LoadWorld(filename string, g *graphics.Graphics) *World {
	wallSprite := sprite.New("resources/block.gif", g)
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
				newTile := tile{
					row:  int32(row),
					col:  int32(col),
					kind: Wall,
					spr:  wallSprite,
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
