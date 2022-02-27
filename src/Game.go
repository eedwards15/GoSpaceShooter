package src

import (
	"SpaceShooter/src/scenes"
	"SpaceShooter/src/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	SceneManager *systems.SceneManager
	keys         []ebiten.Key
}

func (gameClass *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 800
}

func NewGame() *Game {
	g := &Game{}
	g.init()
	return g
}

func (gameClass *Game) init() {
	gameClass.SceneManager = systems.NewSceneManager()
	gameClass.SceneManager.Push(scenes.NewMainMenu())
	gameClass.SceneManager.CurrentScene.Init()
}

func (gameClass *Game) Update() error {
	gameClass.SceneManager.CurrentScene.Update()
	return nil
}

func (gameClass *Game) Draw(screen *ebiten.Image) {
	gameClass.SceneManager.CurrentScene.Draw(screen)
}
