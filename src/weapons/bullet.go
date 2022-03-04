package weapons

import "github.com/hajimehoshi/ebiten/v2"

type Bullet struct {
	Xpos   float64
	Ypos   float64
	Width  float64
	Height float64
	Sprite *ebiten.Image
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

func (bulletClass *Bullet) SetCoordinates(xpos, ypos float64) *Bullet {
	bulletClass.Xpos = xpos
	bulletClass.Ypos = ypos
	return bulletClass
}
