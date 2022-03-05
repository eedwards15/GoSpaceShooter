package npcs

import (
	"SpaceShooter/src/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	PosX  float64
	PosY  float64
	Image *ebiten.Image
}

func NewEnemy() *Enemy {
	e := Enemy{}
	e.Image = systems.ASSETSYSTEM.Assets["Global"].Images["WeakEnemy"]
	e.PosX = float64(systems.WINDOWMANAGER.SCREENWIDTH / 2.0)
	e.PosY = 0
	return &e
}
