package npcs

import (
	"SpaceShooter/src/definitions"
	"SpaceShooter/src/systems"
	"SpaceShooter/src/weapons"
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

type EnemyBase struct {
	posX          float64
	posY          float64
	image         *ebiten.Image
	isDead        bool
	width         int
	height        int
	canShoot      bool
	life          int
	scoreAmount   int
	lastFire      time.Time
	fireSpeed     int64
	movementSpeed float64
}

func NewEnemy(posX, PosY float64, config definitions.EnemyConfig) IEnemy {
	img := systems.ASSETSYSTEM.Assets[config.Assets[0].LocationKey].Images[config.Assets[0].ImageKey]
	w, h := img.Size()
	e := &EnemyBase{
		posY:          PosY,
		posX:          posX,
		width:         w,
		height:        h,
		isDead:        false,
		canShoot:      config.CanShoot,
		life:          config.Life,
		scoreAmount:   config.ScoreAmount,
		fireSpeed:     config.FireRate,
		movementSpeed: config.MovementSpeed,
		lastFire:      time.Now(),
		image:         img,
	}
	//fmt.Printf("%+v\n", e)
	return e
}

func (e *EnemyBase) TakeDamage() {
	e.life = e.life - 1

	if e.life <= 0 {
		e.isDead = true
	}
}

func (e *EnemyBase) GetScoreAmount() int {
	return e.scoreAmount
}

func (e *EnemyBase) Fire() *weapons.Bullet {
	if e.fireSpeed == 0 {
		return nil
	}

	if time.Now().Sub(e.lastFire).Milliseconds() > e.fireSpeed {
		b := weapons.NewBullet(systems.ASSETSYSTEM.Assets["Global"].Images["LaserRed"])
		e.lastFire = time.Now()
		return b
	}

	return nil
}

func (e *EnemyBase) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(e.posX, e.posY)
	screen.DrawImage(e.image, op)
}

func (e *EnemyBase) Update() {
	e.posY += e.movementSpeed
	//Moves the WeakEnemy Back To the Top of the screen
	if e.posY > float64(systems.WINDOWMANAGER.SCREENHEIGHT) {
		e.posY = 0
	}
}

func (e *EnemyBase) SetPosX(x float64) {
	e.posX = x
}

func (e *EnemyBase) SetPosY(y float64) {
	e.posY = y
}

func (e *EnemyBase) GetPosX() float64 {
	return e.posX
}

func (e *EnemyBase) GetPosY() float64 {
	return e.posY
}

func (e *EnemyBase) GetImage() *ebiten.Image {
	return e.image
}

func (e *EnemyBase) IsDead() bool {
	return e.isDead
}

func (e *EnemyBase) Kill() {
	e.isDead = true
}

func (e *EnemyBase) GetWidth() int {
	return e.width
}

func (e *EnemyBase) GetHeight() int {
	return e.height
}

func (e *EnemyBase) CanShoot() bool {
	return e.canShoot
}
