package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"log"
)

type Ship struct {
	tileWidth         int
	tileHeight        int
	row               int
	col               int
	ShipTileSheets    *ebiten.Image
	minX              int
	maxX              int
	minY              int
	maxY              int
	CurrentShipImage  *ebiten.Image
	CurrentShipWidth  float64
	CurrentShipHeight float64
	FireSound         *mp3.Stream
	FireRate          int64
}

func NewShip() *Ship {
	ship := &Ship{}
	ship.tileWidth = 98
	ship.tileHeight = 75
	ship.LoadImages()
	ship.FireRate = 350
	//Move this
	f, _ := ebitenutil.OpenFile("assets/sound effects/414885__matrixxx__retro-laser-shot-03.mp3")
	ship.FireSound, _ = mp3.DecodeWithSampleRate(44100, f)

	return ship
}

func (ship *Ship) LoadImages() {
	img, _, err := ebitenutil.NewImageFromFile("assets/art/player/ships.png")
	if err != nil {
		log.Fatal(err)
	}
	ship.ShipTileSheets = img
}

func (ship *Ship) SelectShip(row int, col int) {
	ship.col = col
	ship.row = row
	ship.minX = (ship.col - 1) * ship.tileWidth
	ship.maxX = (ship.col) * ship.tileWidth
	ship.minY = (ship.row - 1) * ship.tileHeight
	ship.maxY = (ship.row) * ship.tileHeight

	ship.CurrentShipImage = ship.ShipTileSheets.SubImage(image.Rect(ship.minX, ship.minY, ship.maxX, ship.maxY)).(*ebiten.Image)
	width, height := ship.CurrentShipImage.Size()
	ship.CurrentShipWidth = float64(width)
	ship.CurrentShipHeight = float64(height)
}
