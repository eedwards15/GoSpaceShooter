package scenes

import (
	"SpaceShooter/src/player"
	"SpaceShooter/src/systems"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/keyboard/keyboard"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	_ "image/png"
)

type LevelOne struct {
	keys []ebiten.Key
}

func (levelOneClass *LevelOne) Init() {
	player.PLAYER.Ship.SelectShip(1, 2)
}

func (levelOneClass *LevelOne) GetName() string {
	return "Level 1"
}

func NewLevelOne() *LevelOne {
	g := &LevelOne{}
	return g
}

func (levelOneClass *LevelOne) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x00, 0x40, 0x80, 0xff})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(player.PLAYER.XPos, player.PLAYER.YPos)
	screen.DrawImage(player.PLAYER.Ship.CurrentShipImage, op)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (levelOneClass *LevelOne) Update() error {
	levelOneClass.keys = inpututil.AppendPressedKeys(levelOneClass.keys[:0])
	for _, p := range levelOneClass.keys {
		_, ok := keyboard.KeyRect(p)

		if p.String() == "A" {
			player.PLAYER.XPos -= 15.0
		}

		if p.String() == "D" {
			player.PLAYER.XPos += 15
		}

		if p.String() == "W" {
			player.PLAYER.YPos -= 15
		}

		if p.String() == "S" {
			player.PLAYER.YPos += 15
		}

		if p.String() == "Escape" {
			systems.SCENEMANAGER.Pop()
		}

		if !ok {
			continue
		}

	}

	return nil
}
