package main

import (
	"SpaceShooter/src"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	game := src.NewGame()
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Space Shooter")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
