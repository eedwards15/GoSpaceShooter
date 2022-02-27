package scenes

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type LevelOne struct {
}

func (l LevelOne) GetName() string {
	return "Level 1"
}

func NewLevelOne() *LevelOne {
	g := &LevelOne{}
	g.Init()
	return g
}

func (l LevelOne) Init() {
	fmt.Println("Level 1")
}

func (l LevelOne) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x00, 0x40, 0x80, 0xff})
}

func (l LevelOne) Update() error {

	return nil
}
