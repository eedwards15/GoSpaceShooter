package interfaces

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type IScene interface {
	GetName() string
	Init()
	Draw(screen *ebiten.Image)
	Update() error
}
