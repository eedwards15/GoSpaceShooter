package npcs

import (
	"SpaceShooter/src/systems"
	"math/rand"
	"time"
)

type EnemySpawner struct {
	LAST_SPAWN_TIME time.Time
	coolDown        float64
}

func NewEnemySpawner() *EnemySpawner {
	eS := &EnemySpawner{}
	eS.LAST_SPAWN_TIME = time.Now()
	eS.coolDown = 2
	return eS
}

func (enemySpawner *EnemySpawner) SpawnNewEnemy() IEnemy {
	if time.Now().Sub(enemySpawner.LAST_SPAWN_TIME).Seconds() > enemySpawner.coolDown {
		enemySpawner.LAST_SPAWN_TIME = time.Now()
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		xPos := float64(r1.Intn(systems.WINDOWMANAGER.SCREENWIDTH - 100))
		yPos := float64(0)

		s2 := rand.NewSource(time.Now().UnixNano())
		r2 := rand.New(s2)
		x := r2.Intn(101)

		if x < 5 {
			enemySpawner.coolDown = 4
			return NewLevel5Enemy(xPos, yPos)
		}

		if x < 10 {
			enemySpawner.coolDown = 3
			return NewMidRankEnemy(xPos, yPos)

		}

		if x < 30 {
			enemySpawner.coolDown = 1
			return NewWeakEnemy(xPos, yPos)
		}

		if x < 40 {
			enemySpawner.coolDown = 2
			return NewLevel2Enemy(xPos, yPos)

		}

		if x < 50 {
			enemySpawner.coolDown = 3
			return NewLevel3Enemy(xPos, yPos)
		}

	}
	return nil
}
