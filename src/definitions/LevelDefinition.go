package definitions

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
	"path"
)

type LevelDefinition struct {
	BackgroundMusic BackgroundMusic
	Images          map[string]*ebiten.Image
	SoundEffects    map[string]*mp3.Stream
}

func NewLevelDefinition(assetConfig *AssetConfig) *LevelDefinition {
	l := LevelDefinition{}

	//Background Music
	if assetConfig.BackgroundMusic != (BackgroundMusic{}) {
		l.BackgroundMusic = BackgroundMusic{
			Path:       assetConfig.BackgroundMusic.Path,
			SampleRate: assetConfig.BackgroundMusic.SampleRate,
		}
	}

	//Images
	l.Images = make(map[string]*ebiten.Image)
	for i := 0; i < len(assetConfig.Images); i++ {
		record := assetConfig.Images[i]
		path := path.Join("assets", record.Location, record.FileName)
		l.Images[record.Key] = openImage(path)
	}

	//Sound Effects
	l.SoundEffects = make(map[string]*mp3.Stream)
	for i := 0; i < len(assetConfig.SoundEffects); i++ {
		record := assetConfig.SoundEffects[i]
		path := path.Join("assets", record.Location, record.FileName)
		l.SoundEffects[record.Key] = openSound(path, record.SampleRate)
	}

	return &l
}

type BackgroundMusic struct {
	Path       string
	SampleRate int
}

func openImage(location string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(location)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
func openSound(location string, sampleRate int) *mp3.Stream {
	f, _ := ebitenutil.OpenFile(location)
	fireSound, _ := mp3.DecodeWithSampleRate(sampleRate, f)
	return fireSound
}
