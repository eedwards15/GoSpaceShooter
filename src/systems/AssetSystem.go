package systems

import (
	"SpaceShooter/src/definitions"
	"SpaceShooter/src/helpers"
	"sync"
)

var (
	ASSETSYSTEM *AssetSystem
)

type AssetSystem struct {
	Assets map[string]*definitions.LevelDefinition
}

func InitAssetSystem() {
	ASSETSYSTEM = &AssetSystem{}
	ASSETSYSTEM.Assets = make(map[string]*definitions.LevelDefinition)
	configValues, _ := helpers.AssetConfigHelper()

	var wg sync.WaitGroup
	for i := 0; i < len(configValues); i++ {
		wg.Add(1)
		r := configValues[i]
		go func(record *definitions.AssetConfig) {
			defer wg.Done()
			ASSETSYSTEM.Assets[record.Scene] = definitions.NewLevelDefinition(record)

		}(r)
	}
	wg.Wait()

}
