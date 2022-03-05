package helpers

import (
	"SpaceShooter/src/definitions"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func AssetConfigHelper() ([]*definitions.AssetConfig, error) {
	staticPath := "assets/configs"
	files, err := ioutil.ReadDir(staticPath)
	if err != nil {
		log.Fatal(err)
	}

	assetConfigs := []*definitions.AssetConfig{}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileValue, _ := os.ReadFile(path.Join(staticPath, file.Name()))
		assetConfig := definitions.AssetConfig{}
		json.Unmarshal(fileValue, &assetConfig)
		assetConfigs = append(assetConfigs, &assetConfig)
	}

	return assetConfigs, nil
}
