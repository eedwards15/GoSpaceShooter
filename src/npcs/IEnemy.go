package npcs

import (
	"SpaceShooter/src/weapons"
	"github.com/hajimehoshi/ebiten/v2"
)

type IEnemy interface {
	SetPosX(x float64)
	SetPosY(y float64)
	GetPosX() float64
	GetPosY() float64
	GetImage() *ebiten.Image
	IsDead() bool
	TakeDamage()
	GetWidth() int
	GetHeight() int
	CanShoot() bool
	GetScoreAmount() int
	Fire() *weapons.Bullet
	Draw(screen *ebiten.Image)
	Update()
}
