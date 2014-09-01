package game

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"

	"github.com/veandco/go-sdl2/sdl"

	"graphics"
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

var Empty tileKind = tileKind{v: 0}
var Wall tileKind = tileKind{v: 1}
var PlayerStart tileKind = tileKind{v: 2}
var MoblinStart tileKind = tileKind{v: 3}

type tile struct {
	row  int32
	col  int32
	kind tileKind
	spr  *graphics.Sprite
}

type World struct {
	Player  Player
	Enemies []Slime
	tiles   []tile
}

var wallSprite *graphics.Sprite

func LoadWorld(filename string, g *graphics.Graphics) *World {
	wallSprite = graphics.NewSprite(3, 0, 1, 1, g)

	f, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Unable to open world file: %s", filename))
	}
	defer f.Close()

	player := NewPlayer(g)
	enemies := make([]Slime, 0)
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
			case '@':
				newTile := tile{
					row:  int32(row),
					col:  int32(col),
					kind: PlayerStart,
				}
				tiles = append(tiles, newTile)

				bounds := newTile.Bounds()
				player.x = float32(bounds.X)
				player.y = float32(bounds.Y)

			case 'm':
				newTile := tile{
					row:  int32(row),
					col:  int32(col),
					kind: MoblinStart,
				}
				tiles = append(tiles, newTile)

				slime := NewSlime(g)
				bounds := newTile.Bounds()
				slime.x = float32(bounds.X)
				slime.y = float32(bounds.Y)
				slime.goal = &tile{
					row: newTile.row + rand.Int31n(10) - 5,
					col: newTile.col + rand.Int31n(10) - 5,
				}
				enemies = append(enemies, *slime)
			}
		}

		row += 1
		line, isPrefix, err = r.ReadLine()
	}

	return &World{
		Player:  *player,
		Enemies: enemies,
		tiles:   tiles,
	}
}

func (w *World) CollideWithTiles(b Bounded, d Direction) bool {
	collided := false
	for _, t := range w.tiles {
		if t.IsPassable() {
			continue
		}
		bRect := b.Bounds()
		tRect := t.Bounds()
		if bRect.HasIntersection(tRect) {
			collided = true
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
	return collided
}

func (w *World) FindTileKind(tk tileKind) *tile {
	for _, t := range w.tiles {
		if t.kind == tk {
			return &t
		}
	}
	return nil
}

func (w *World) TileAt(row int32, col int32) *tile {
	// First, search for a tile
	for _, tile := range w.tiles {
		if tile.row == row && tile.col == col {
			return &tile
		}
	}

	// If not found, create a new "Empty" tile and return it.
	return &tile{
		row: row,
		col: col,
	}
}

func (w *World) Update(dt uint32) {
	w.Player.Update(dt, w)
	for i := range w.Enemies {
		w.Enemies[i].Update(dt, w)
	}

	// Remove dead enemies, drop loot
	i := 0
	for i < len(w.Enemies) {
		if w.Enemies[i].health <= 0 {
			w.Enemies[i], w.Enemies = w.Enemies[len(w.Enemies)-1], w.Enemies[:len(w.Enemies)-1]
		} else {
			i += 1
		}
	}
}

func (w *World) Draw() {
	w.Player.Draw()
	for _, enemy := range w.Enemies {
		enemy.Draw()
	}
	for _, tile := range w.tiles {
		tile.Draw()
	}
}

func (t *tile) IsPassable() bool {
	switch t.kind {
	case Wall:
		return false
	default:
		return true
	}
}

func (t *tile) Draw() {
	if t.spr == nil {
		return
	}
	t.spr.X = float32(uint(t.col) * t.spr.W)
	t.spr.Y = float32(uint(t.row) * t.spr.H)
	t.spr.Draw()
}

func (t *tile) Bounds() *sdl.Rect {
	r := &sdl.Rect{t.col * int32(wallSprite.W), t.row * int32(wallSprite.H), 0, 0}
	if t.spr != nil {
		r.W = int32(t.spr.W)
		r.H = int32(t.spr.H)
	}
	return r
}
