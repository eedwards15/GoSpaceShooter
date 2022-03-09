package npcs

import "SpaceShooter/src/systems"

type Level3Enemy struct {
	EnemyBase
}

func NewLevel3Enemy(x, y float64) IEnemy {
	img := systems.ASSETSYSTEM.Assets["Global"].Images["Level3Enemy"]
	w, h := img.Size()
	e := &Level3Enemy{
		EnemyBase{
			image:         img,
			isDead:        false,
			width:         w,
			height:        h,
			canShoot:      true,
			posX:          x,
			posY:          y,
			life:          4,
			scoreAmount:   100,
			fireSpeed:     1000,
			movementSpeed: 2,
		},
	}

	return e
}

func (e *Level3Enemy) Update() {
	e.posY += 2
	//Moves the WeakEnemy Back To the Top of the screen
	if e.posY > float64(systems.WINDOWMANAGER.SCREENHEIGHT) {
		e.posY = 0
	}
}
