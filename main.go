package main

import (
	"SpaceShooter/src"
	"SpaceShooter/src/player"
	"SpaceShooter/src/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	systems.SCENEMANAGER = systems.NewSceneManager()
	systems.WINDOWMANAGER = systems.WindowManager{SCREENHEIGHT: 800, SCREENWIDTH: 1280}
	player.NewPLayer()

	game := src.NewGame()
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Space Shooter")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
func GetSceneManager() *systems.SceneManager {
	return systems.NewSceneManager()
}
