package scenes

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/keyboard/keyboard"
	"image/color"
)

type LevelOne struct {
	keys []ebiten.Key
}

func (levelOneClass *LevelOne) Init() {
	//levelOneClass.sceneManager = sceneManager
	fmt.Println("Level 1")
}

func (levelOneClass *LevelOne) GetName() string {
	return "Level 1"
}

func NewLevelOne() *LevelOne {
	g := &LevelOne{}
	return g
}

func (levelOneClass *LevelOne) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x00, 0x40, 0x80, 0xff})
}

func (levelOneClass *LevelOne) Update() error {
	for _, p := range levelOneClass.keys {
		_, ok := keyboard.KeyRect(p)

		if p.String() == "Escape" {
			//levelOneClass.sceneManager.SceneManager.Pop()
		}

		if !ok {
			continue
		}

	}

	return nil
}
