package player

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Ship *Ship
	XPos float64
	YPos float64
}

func NewPLayer(x, y float64) *Player {
	PLAYER := &Player{}
	PLAYER.XPos = x
	PLAYER.YPos = y
	PLAYER.Ship = NewShip()
	return PLAYER
}

func (player *Player) MoveX(x float64) {
	player.XPos += x
}

func (player *Player) MoveY(y float64) {
	player.YPos += y
}

func (player *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterLinear
	op.GeoM.Translate(player.XPos, player.YPos)
	screen.DrawImage(player.Ship.CurrentShipImage, op)

}
