package npcs

import "SpaceShooter/src/systems"

type Level5Enemy struct {
	EnemyBase
}

func NewLevel5Enemy(x, y float64) IEnemy {
	img := systems.ASSETSYSTEM.Assets["Global"].Images["Level5Enemy"]
	w, h := img.Size()
	e := &Level5Enemy{
		EnemyBase{
			image:         img,
			isDead:        false,
			width:         w,
			height:        h,
			canShoot:      true,
			posX:          x,
			posY:          y,
			life:          6,
			scoreAmount:   600,
			fireSpeed:     1400,
			movementSpeed: 3,
		},
	}

	return e
}
