package definitions

import (
	"SpaceShooter/src/helpers"
	"github.com/hajimehoshi/ebiten/v2"
)

type LevelDefinition struct {
	BackgroundMusic BackgroundMusic
	Images          map[string]*ebiten.Image
}

func NewLevelDefinition() *LevelDefinition {
	l := LevelDefinition{}
	l.Images = make(map[string]*ebiten.Image)
	return &l
}

func (LevelDefinitionClass *LevelDefinition) AddImage(key string, location string) {
	LevelDefinitionClass.Images[key] = helpers.OpenImage(location)
}

type BackgroundMusic struct {
	Path       string
	SampleRate int
}
