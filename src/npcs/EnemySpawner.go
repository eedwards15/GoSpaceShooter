package npcs

import (
	"SpaceShooter/assets"
	"SpaceShooter/src/definitions"
	"SpaceShooter/src/systems"
	"encoding/json"
	"math/rand"
	"path"
	"time"
)

type EnemySpawner struct {
	enemyConfigs  []definitions.EnemyConfig
	LastSpawnTime time.Time
	coolDown      float64
}

func NewEnemySpawner() *EnemySpawner {
	eS := &EnemySpawner{}
	eS.LastSpawnTime = time.Now()
	eS.coolDown = 2
	configs, _ := assets.AssetsFileSystem.ReadDir("settings/enemy")
	for i := 0; i < len(configs); i++ {
		fileValue, _ := assets.AssetsFileSystem.ReadFile(path.Join("settings/enemy", configs[i].Name()))
		enemyConfig := definitions.EnemyConfig{}
		json.Unmarshal(fileValue, &enemyConfig)
		eS.enemyConfigs = append(eS.enemyConfigs, enemyConfig)
	}
	return eS
}

func (enemySpawner *EnemySpawner) SpawnNewEnemy() IEnemy {
	if time.Now().Sub(enemySpawner.LastSpawnTime).Seconds() > enemySpawner.coolDown {
		enemySpawner.LastSpawnTime = time.Now()
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		xPos := float64(r1.Intn(systems.WINDOWMANAGER.SCREENWIDTH - 100))
		yPos := float64(0)

		s2 := rand.NewSource(time.Now().UnixNano())
		r2 := rand.New(s2)
		x := r2.Intn(len(enemySpawner.enemyConfigs))

		enemySpawner.coolDown = enemySpawner.enemyConfigs[x].CoolDown
		e := NewEnemy(xPos, yPos, enemySpawner.enemyConfigs[x])
		return e

	}
	return nil
}
