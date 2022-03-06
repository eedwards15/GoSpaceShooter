package npcs

import (
	"SpaceShooter/src/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

type WeakEnemy struct {
	posX     float64
	posY     float64
	image    *ebiten.Image
	isDead   bool
	width    int
	height   int
	canShoot bool
	life     int
}

func NewWeakEnemy(x, y float64) IEnemy {
	img := systems.ASSETSYSTEM.Assets["Global"].Images["WeakEnemy"]
	w, h := img.Size()

	e := &WeakEnemy{
		image:    img,
		isDead:   false,
		width:    w,
		height:   h,
		canShoot: false,
		posX:     x,
		posY:     y,
		life:     1,
	}

	return e
}

func (e *WeakEnemy) SetPosX(x float64) {
	e.posX = x
}

func (e *WeakEnemy) SetPosY(y float64) {
	e.posY = y
}

func (e WeakEnemy) GetPosX() float64 {
	return e.posX
}

func (e WeakEnemy) GetPosY() float64 {
	return e.posY
}

func (e WeakEnemy) GetImage() *ebiten.Image {
	return e.image
}

func (e WeakEnemy) IsDead() bool {
	return e.isDead
}

func (e *WeakEnemy) TakeDamage() {
	e.life = e.life - 1

	if e.life <= 0 {
		e.isDead = true
	}
}

func (e WeakEnemy) GetWidth() int {
	return e.width
}

func (e WeakEnemy) GetHeight() int {
	return e.height
}

func (e WeakEnemy) CanShoot() bool {
	return e.canShoot
}
