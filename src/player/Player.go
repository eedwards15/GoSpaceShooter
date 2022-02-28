package player

var (
	PLAYER *Player
)

type Player struct {
	Ship *Ship
	XPos float64
	YPos float64
}

func NewPLayer() {
	if PLAYER == nil {
		PLAYER = &Player{}
		PLAYER.Ship = NewShip()
	}
}
