package main

import (
	"SpaceShooter/src"
	"SpaceShooter/src/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	systems.InitAssetSystem()
	systems.InitSceneManager()
	systems.InitWindowManager(1280, 800)
	systems.MUSICSYSTEM = systems.NewMusicSystem(systems.ASSETSYSTEM.Assets["MainMenu"].BackgroundMusic) //Refactor how this works

	game := src.NewGame()
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Space Shooter")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
