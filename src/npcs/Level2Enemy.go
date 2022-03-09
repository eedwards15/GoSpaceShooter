package npcs

import "SpaceShooter/src/systems"

type Level2Enemy struct {
	EnemyBase
}

func NewLevel2Enemy(x, y float64) IEnemy {
	img := systems.ASSETSYSTEM.Assets["Global"].Images["Level2Enemy"]
	w, h := img.Size()

	e := &Level2Enemy{
		EnemyBase{
			image:         img,
			isDead:        false,
			width:         w,
			height:        h,
			canShoot:      true,
			posX:          x,
			posY:          y,
			life:          2,
			scoreAmount:   35,
			fireSpeed:     1200,
			movementSpeed: 5,
		},
	}

	return e
}

func (e *Level2Enemy) Update() {
	e.posY += 5
	//Moves the WeakEnemy Back To the Top of the screen
	if e.posY > float64(systems.WINDOWMANAGER.SCREENHEIGHT) {
		e.posY = 0
	}
}
