package npcs

import (
	"SpaceShooter/src/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	PosX   float64
	PosY   float64
	Image  *ebiten.Image
	Dead   bool
	Width  int
	Height int
}

func NewEnemy(x, y float64) *Enemy {
	e := Enemy{}
	e.Dead = false
	e.Image = systems.ASSETSYSTEM.Assets["Global"].Images["WeakEnemy"]
	w, h := e.Image.Size()
	e.Width = w
	e.Height = h
	e.PosX = x
	e.PosY = y
	return &e
}
