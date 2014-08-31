package game

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"

	"graphics"
	"sprite"
)

type Moblin struct {
	x         float32
	y         float32
	speed     float32 // pixels/s
	health    int32
	maxHealth int32
	direction Direction
	goal      *tile
	spr       *sprite.Sprite
}

func NewMoblin(g *graphics.Graphics) *Moblin {
	return &Moblin{
		health:    100,
		maxHealth: 100,
		speed:     50,
		spr:       sprite.New("resources/moblin.gif", g),
	}
}

// TODO: Goal should be generated from current position, not current goal.
func (m *Moblin) RandomGoal(w *World) *tile {
	newRow := m.goal.row + rand.Int31n(10) - 5
	newCol := m.goal.col + rand.Int31n(10) - 5
	return w.TileAt(newRow, newCol)
}

func (m *Moblin) DirectionToGoal(w *World) Direction {
	if m.goal == nil {
		return 0
	}
	gRect := m.goal.Bounds()
	xDist := int32(m.x) - gRect.X
	yDist := int32(m.y) - gRect.Y
	xAxis := true
	if math.Abs(float64(yDist)) > math.Abs(float64(xDist)) {
		xAxis = false
	}

	switch {
	case xAxis && xDist < 0:
		return East
	case xAxis && xDist > 0:
		return West
	case !xAxis && yDist < 0:
		return South
	case !xAxis && yDist > 0:
		return North
	default:
		// Goal reached, so create a new one.
		m.goal = m.RandomGoal(w)
		return m.DirectionToGoal(w)
	}
}

func (m *Moblin) Update(dt uint32, w *World) {
	m.direction = m.DirectionToGoal(w)
	velocity := m.speed * float32(dt) / 1000
	collided := false
	if m.direction&North > 0 {
		m.y -= velocity
		if w.CollideWithTiles(m, North) {
			collided = true
		}
	}
	if m.direction&East > 0 {
		m.x += velocity
		if w.CollideWithTiles(m, East) {
			collided = true
		}
	}
	if m.direction&South > 0 {
		m.y += velocity
		if w.CollideWithTiles(m, South) {
			collided = true
		}
	}
	if m.direction&West > 0 {
		m.x -= velocity
		if w.CollideWithTiles(m, West) {
			collided = true
		}
	}
	if collided {
		m.goal = m.RandomGoal(w)
	}
	fmt.Println(m.goal)
}

func (m *Moblin) ChangeHealth(amount int32) {
	m.health += amount
	fmt.Printf("Moblin Health: %d/%d\n", m.health, m.maxHealth)
}

func (m *Moblin) Bounds() *sdl.Rect {
	return &sdl.Rect{int32(m.x), int32(m.y), int32(m.spr.W), int32(m.spr.H)}
}

func (m *Moblin) SetBounds(r *sdl.Rect) {
	m.x = float32(r.X)
	m.y = float32(r.Y)
}

func (m *Moblin) Draw() {
	m.spr.X = m.x
	m.spr.Y = m.y
	m.spr.Draw()
}
