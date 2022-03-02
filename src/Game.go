package src

import (
	"SpaceShooter/src/scenes"
	"SpaceShooter/src/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	keys []ebiten.Key
}

func (gameClass *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return systems.WINDOWMANAGER.SCREENWIDTH, systems.WINDOWMANAGER.SCREENHEIGHT
}

func NewGame() *Game {
	g := &Game{}
	g.init()
	return g
}

func (gameClass *Game) init() {
	systems.MUSICSYSTEM.PlaySong()
	systems.SCENEMANAGER.Push(scenes.NewMainMenu())
}

func (gameClass *Game) Update() error {
	systems.SCENEMANAGER.CurrentScene.Update()
	return nil
}

func (gameClass *Game) Draw(screen *ebiten.Image) {
	systems.SCENEMANAGER.CurrentScene.Draw(screen)
}
