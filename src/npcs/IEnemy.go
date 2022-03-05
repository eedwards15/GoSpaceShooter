package npcs

import "github.com/hajimehoshi/ebiten/v2"

type IEnemy interface {
	SetPosX() float64
	SetPosY() float64
	GetImage() *ebiten.Image
	IsDead() bool
	Kill()
	GetWidth() int
	GetHeight() int
	CanShoot() bool
}
