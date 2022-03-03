package scenes

import (
	"SpaceShooter/src/npcs"
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

//Convert this into something that loads the levels.

type Level struct {
	keys  []ebiten.Key
	enemy *npcs.Enemy
}

func (levelOneClass *Level) Init() {
	systems.MUSICSYSTEM.LoadSong(systems.ASSETSYSTEM.LevelOne.BackgroundMusic).PlaySong()
	player.PLAYER.Ship.SelectShip(1, 2)
	levelOneClass.enemy = npcs.NewEnemy()

	cX, cY := systems.WINDOWMANAGER.Center()
	player.PLAYER.XPos = cX
	player.PLAYER.YPos = cY

}

func (levelOneClass *Level) GetName() string {
	return "Level 1"
}

func NewLevelOne() *Level {
	g := &Level{}
	return g
}

func (levelOneClass *Level) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x00, 0x40, 0x80, 0xff})

	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterLinear
	op.GeoM.Translate(player.PLAYER.XPos, player.PLAYER.YPos)
	screen.DrawImage(player.PLAYER.Ship.CurrentShipImage, op)

	op2 := &ebiten.DrawImageOptions{}
	op2.Filter = ebiten.FilterLinear
	op2.GeoM.Translate(levelOneClass.enemy.PosX, levelOneClass.enemy.PosY)
	screen.DrawImage(levelOneClass.enemy.Image, op2)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (levelOneClass *Level) Update() error {

	levelOneClass.keys = inpututil.AppendPressedKeys(levelOneClass.keys[:0])
	for _, p := range levelOneClass.keys {
		_, ok := keyboard.KeyRect(p)

		if p.String() == "A" && (player.PLAYER.XPos > 0) {
			player.PLAYER.XPos -= 10
		}

		if p.String() == "D" && (player.PLAYER.XPos+(player.PLAYER.Ship.CurrentShipWidth) < float64(systems.WINDOWMANAGER.SCREENWIDTH)) {
			player.PLAYER.XPos += 10
		}

		if p.String() == "W" && (player.PLAYER.YPos > 0) {
			player.PLAYER.YPos -= 10
		}

		if p.String() == "S" && ((player.PLAYER.YPos + player.PLAYER.Ship.CurrentShipHeight) < float64(systems.WINDOWMANAGER.SCREENHEIGHT)) {
			player.PLAYER.YPos += 10
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
