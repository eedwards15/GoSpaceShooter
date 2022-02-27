package src

import (
	"SpaceShooter/src/interfaces"
	"SpaceShooter/src/scenes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/keyboard/keyboard"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	currentScene interfaces.IScene
	allScenese   scenes.SceneManager
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
	gameClass.allScenese = gameClass.allScenese.Push(scenes.NewMainMenu())
	gameClass.currentScene = *gameClass.allScenese.Peek()

}

func (gameClass *Game) Update() error {

	gameClass.keys = inpututil.AppendPressedKeys(gameClass.keys[:0])

	for _, p := range gameClass.keys {
		_, ok := keyboard.KeyRect(p)

		println("Key", p.String())

		if p.String() == "S" && gameClass.currentScene.GetName() != "Main Menu" {
			if len(gameClass.allScenese) > 1 {
				gameClass.allScenese, _ = gameClass.allScenese.Pop()
			}
			gameClass.currentScene = *gameClass.allScenese.Peek()

		}

		if p.String() == "A" && gameClass.currentScene.GetName() != "Level 1" {
			gameClass.allScenese = gameClass.allScenese.Push(scenes.NewLevelOne())
			gameClass.currentScene = *gameClass.allScenese.Peek()
			gameClass.currentScene.Init()
		}

		if !ok {
			continue
		}

	}

	gameClass.currentScene.Update()
	return nil
}

func (gameClass *Game) Draw(screen *ebiten.Image) {
	gameClass.currentScene.Draw(screen)
}
