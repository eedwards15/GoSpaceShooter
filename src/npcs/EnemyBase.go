package npcs

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type EnemyBase struct {
	posX     float64
	posY     float64
	image    *ebiten.Image
	isDead   bool
	width    int
	height   int
	canShoot bool
}

func (e *EnemyBase) SetPosX(x float64) {
	e.posX = x
}

func (e *EnemyBase) SetPosY(y float64) {
	e.posY = y
}

func (e EnemyBase) GetPosX() float64 {
	return e.posX
}

func (e EnemyBase) GetPosY() float64 {
	return e.posY
}

func (e EnemyBase) GetImage() *ebiten.Image {
	return e.image
}

func (e EnemyBase) IsDead() bool {
	return e.isDead
}

func (e *EnemyBase) Kill() {
	e.isDead = true
}

func (e EnemyBase) GetWidth() int {
	return e.width
}

func (e EnemyBase) GetHeight() int {
	return e.height
}

func (e EnemyBase) CanShoot() bool {
	return e.canShoot
}

//
//func NewEnemy(x, y float64) *WeakEnemy {
//	e := WeakEnemy{}
//	e.Dead = false
//	e.Image = systems.ASSETSYSTEM.Assets["Global"].Images["WeakEnemy"]
//	w, h := e.Image.Size()
//	e.Width = w
//	e.Height = h
//	e.PosX = x
//	e.PosY = y
//	return &e
//}
