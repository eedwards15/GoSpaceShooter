package player

type Player struct {
	Ship *Ship
	XPos float64
	YPos float64
}

func NewPLayer() *Player {
	PLAYER := &Player{}
	PLAYER.Ship = NewShip()
	return PLAYER
}
