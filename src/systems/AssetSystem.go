package systems

import "SpaceShooter/src/definitions"

var (
	ASSETSYSTEM *AssetSystem
)

//Make This Load From Config Files
type AssetSystem struct {
	MainMenu definitions.LevelDefinition
	LevelOne definitions.LevelDefinition
}

func NewAssetSystem() *AssetSystem {
	assetsystem := AssetSystem{}
	assetsystem.MainMenu = *definitions.NewLevelDefinition()
	assetsystem.MainMenu.BackgroundMusic = definitions.BackgroundMusic{
		Path:       "assets/music/545452__bertsz__organ-type-hiphop-beat.mp3",
		SampleRate: 44100,
	}

	assetsystem.LevelOne = *definitions.NewLevelDefinition()
	assetsystem.LevelOne.BackgroundMusic = definitions.BackgroundMusic{
		Path:       "assets/music/416984__soundflakes__desolated-field.mp3",
		SampleRate: 44100,
	}
	assetsystem.LevelOne.AddImage("Background", "assets/art/scene/Space-Background-Tiled.png")
	assetsystem.LevelOne.AddImage("LaserBullet", "assets/art/weapons/laser_bullet.png")

	return &assetsystem
}
