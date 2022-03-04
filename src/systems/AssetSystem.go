package systems

import (
	"SpaceShooter/src/definitions"
	"SpaceShooter/src/helpers"
)

var (
	ASSETSYSTEM *AssetSystem
)

//Make This Load From Config Files
type AssetSystem struct {
	Assets map[string]*definitions.LevelDefinition
}

func NewAssetSystem() *AssetSystem {
	assetsystem := AssetSystem{}
	assetsystem.Assets = make(map[string]*definitions.LevelDefinition)
	configValues, _ := helpers.AssetConfigHelper()

	for i := 0; i < len(configValues); i++ {
		record := configValues[i]
		assetsystem.Assets[record.Scene] = definitions.NewLevelDefinition(record)

	}

	//
	//assetsystem.MainMenu = *definitions.NewLevelDefinition()
	//assetsystem.MainMenu.BackgroundMusic = definitions.BackgroundMusic{
	//	Path:       "assets/music/545452__bertsz__organ-type-hiphop-beat.mp3",
	//	SampleRate: 44100,
	//}
	//
	//assetsystem.LevelOne = *definitions.NewLevelDefinition()
	//assetsystem.LevelOne.BackgroundMusic = definitions.BackgroundMusic{
	//	Path:       "assets/music/416984__soundflakes__desolated-field.mp3",
	//	SampleRate: 44100,
	//}
	//assetsystem.LevelOne.AddImage("Background", "assets/art/scene/Space-Background-Tiled.png")
	//assetsystem.LevelOne.AddImage("LaserBullet", "assets/art/weapons/laser_bullet.png")

	return &assetsystem
}
