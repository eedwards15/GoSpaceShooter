package npcs

import (
	"SpaceShooter/src/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Enemy struct {
	PosX  float64
	PosY  float64
	Image *ebiten.Image
}

//TODO:
//make this function take in a string for file name
func NewEnemy() *Enemy {
	e := Enemy{}
	e.Image = e.loadImages()
	e.PosX = float64(systems.WINDOWMANAGER.SCREENWIDTH / 2.0)
	e.PosY = 0
	return &e
}

//TODO:
//Make this be part of the Asset System
//
func (e Enemy) loadImages() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile("assets/art/enemies/enemyBlack1.png")
	if err != nil {
		log.Fatal(err)
	}
	return img
}
