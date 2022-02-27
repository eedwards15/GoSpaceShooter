package player

type Player struct {
	Ship *Ship
	XPos float64
	YPos float64
}

func NewPLayer() *Player {
	player := &Player{}
	player.Ship = NewShip()
	return player
}
