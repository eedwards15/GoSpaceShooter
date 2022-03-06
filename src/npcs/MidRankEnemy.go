package npcs

import (
	"SpaceShooter/src/systems"
	"SpaceShooter/src/weapons"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

type MidRankEnemy struct {
	posX        float64
	posY        float64
	image       *ebiten.Image
	isDead      bool
	width       int
	height      int
	canShoot    bool
	life        int
	scoreAmount int
	lastFire    time.Time
}

func (e *MidRankEnemy) Fire() *weapons.Bullet {
	if time.Now().Sub(e.lastFire).Milliseconds() > 800 {
		fmt.Println("fire")
		b := weapons.NewBullet(systems.ASSETSYSTEM.Assets["Global"].Images["LaserRed"])
		e.lastFire = time.Now()
		return b
	}
	return nil
}

func NewMidRankEnemy(x, y float64) IEnemy {
	img := systems.ASSETSYSTEM.Assets["Global"].Images["MidRankEnemy"]
	w, h := img.Size()

	e := &MidRankEnemy{
		image:       img,
		isDead:      false,
		width:       w,
		height:      h,
		canShoot:    true,
		posX:        x,
		posY:        y,
		life:        3,
		scoreAmount: 25,
	}

	return e
}
func (e MidRankEnemy) GetScoreAmount() int {
	return e.scoreAmount
}

func (e *MidRankEnemy) SetPosX(x float64) {
	e.posX = x
}

func (e *MidRankEnemy) SetPosY(y float64) {
	e.posY = y
}

func (e MidRankEnemy) GetPosX() float64 {
	return e.posX
}

func (e MidRankEnemy) GetPosY() float64 {
	return e.posY
}

func (e MidRankEnemy) GetImage() *ebiten.Image {
	return e.image
}

func (e MidRankEnemy) IsDead() bool {
	return e.isDead
}

func (e *MidRankEnemy) TakeDamage() {
	e.life = e.life - 1

	if e.life <= 0 {
		e.isDead = true
	}
}

func (e MidRankEnemy) GetWidth() int {
	return e.width
}

func (e MidRankEnemy) GetHeight() int {
	return e.height
}

func (e MidRankEnemy) CanShoot() bool {
	return e.canShoot
}
