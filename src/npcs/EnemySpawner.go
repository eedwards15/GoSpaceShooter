package npcs

import (
	"math/rand"
	"time"
)

type EnemySpawner struct{}

func SpawnNewEnemy(xPos, yPos float64) IEnemy {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	x := r1.Intn(101)

	if x < 90 {
		return NewWeakEnemy(xPos, yPos)
	}

	return NewMidRankEnemy(xPos, yPos)

}
