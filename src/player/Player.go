package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

type Player struct {
	MaxLife      int
	Life         int
	Ship         *Ship
	XPos         float64
	YPos         float64
	IsDead       bool
	invulnerable time.Time
}

func NewPLayer(x, y float64) *Player {

	PLAYER := &Player{}
	PLAYER.XPos = x
	PLAYER.YPos = y
	PLAYER.Ship = NewShip()
	PLAYER.IsDead = false
	PLAYER.Life = 3
	PLAYER.invulnerable = time.Now()
	return PLAYER
}

func (player *Player) MoveX(x float64) {
	player.XPos += x
}

func (player *Player) MoveY(y float64) {
	player.YPos += y
}

func (player *Player) IsInvulnerable() bool {
	return player.invulnerable.Unix() > time.Now().Unix()
}

func (player *Player) TakeDamage() {
	if player.invulnerable.Unix() > time.Now().Unix() {
		return
	}
	player.Life -= 1
	player.invulnerable = time.Now().Add(time.Second * 3)
	if player.Life <= 0 {
		player.IsDead = true
	}
}

func (player *Player) Heal() {
	if player.Life > 0 && player.Life < player.MaxLife {
		player.Life += 1
	}
}

func (player *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterLinear

	if player.IsInvulnerable() {
		op.ColorM.Scale(1, 1, 1, .50)
	} else {
		op.ColorM.Scale(1, 1, 1, 1)
	}

	op.GeoM.Translate(player.XPos, player.YPos)
	screen.DrawImage(player.Ship.CurrentShipImage, op)

}
