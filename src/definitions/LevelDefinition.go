package definitions

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"path"
)

type LevelDefinition struct {
	BackgroundMusic BackgroundMusic
	Images          map[string]*ebiten.Image
}

func NewLevelDefinition(assetConfig *AssetConfig) *LevelDefinition {
	l := LevelDefinition{}
	l.BackgroundMusic = BackgroundMusic{
		Path:       assetConfig.BackgroundMusic.Path,
		SampleRate: assetConfig.BackgroundMusic.SampleRate,
	}

	l.Images = make(map[string]*ebiten.Image)
	for i := 0; i < len(assetConfig.Images); i++ {
		record := assetConfig.Images[i]
		path := path.Join("assets", record.Location, record.FileName)
		l.Images[record.Key] = OpenImage(path)
	}
	return &l
}

type BackgroundMusic struct {
	Path       string
	SampleRate int
}

func OpenImage(location string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(location)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
