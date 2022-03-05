package main

import (
	"SpaceShooter/src"
	"SpaceShooter/src/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	systems.InitAssetSystem()
	systems.SCENEMANAGER = systems.NewSceneManager()
	systems.WINDOWMANAGER = systems.NewWindowManager(1280, 800)
	systems.MUSICSYSTEM = systems.NewMusicSystem(systems.ASSETSYSTEM.Assets["MainMenu"].BackgroundMusic)

	game := src.NewGame()
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Space Shooter")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
