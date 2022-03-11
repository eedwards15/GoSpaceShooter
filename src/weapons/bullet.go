package weapons

import (
	"SpaceShooter/src/models"
	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	Width  float64
	Height float64
	Sprite *ebiten.Image
	models.Vector2
}

func NewBullet(sprite *ebiten.Image) *Bullet {
	w, h := sprite.Size()

	b := Bullet{
		Sprite: sprite,
		Width:  float64(w),
		Height: float64(h),
	}
	return &b
}

func (bulletClass *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(bulletClass.Xpos, bulletClass.Ypos)
	screen.DrawImage(bulletClass.Sprite, op)
}

func (bulletClass *Bullet) SetCoordinates(xpos, ypos float64) *Bullet {
	bulletClass.Xpos = xpos
	bulletClass.Ypos = ypos
	return bulletClass
}
