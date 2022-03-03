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
	systems.WINDOWMANAGER = systems.NewWindowManager(1280, 800)
	systems.ASSETSYSTEM = systems.NewAssetSystem()
	systems.MUSICSYSTEM = systems.NewMusicSystem(systems.ASSETSYSTEM.MainMenu.BackgroundMusic)
	player.NewPLayer()

	game := src.NewGame()
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Space Shooter")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
