package helpers

import (
	"SpaceShooter/assets"
	"SpaceShooter/src/definitions"
	"encoding/json"
	"log"
	"path"
)

func AssetConfigHelper() ([]*definitions.AssetConfig, error) {
	files, err := assets.AssetsFileSystem.ReadDir("configs")
	if err != nil {
		log.Fatal(err)
	}

	assetConfigs := []*definitions.AssetConfig{}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileValue, _ := assets.AssetsFileSystem.ReadFile(path.Join("configs", file.Name()))
		assetConfig := definitions.AssetConfig{}
		json.Unmarshal(fileValue, &assetConfig)
		assetConfigs = append(assetConfigs, &assetConfig)
	}

	return assetConfigs, nil
}
