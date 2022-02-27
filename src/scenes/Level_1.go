package scenes

import (
	"SpaceShooter/src/systems"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/keyboard/keyboard"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
)

type LevelOne struct {
	keys []ebiten.Key
}

func (levelOneClass *LevelOne) Init() {
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
	levelOneClass.keys = inpututil.AppendPressedKeys(levelOneClass.keys[:0])
	for _, p := range levelOneClass.keys {
		_, ok := keyboard.KeyRect(p)

		if p.String() == "Escape" {
			systems.SCENEMANAGER.Pop()
		}

		if !ok {
			continue
		}

	}

	return nil
}
