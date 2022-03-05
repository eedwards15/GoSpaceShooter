package scenes

import (
	"SpaceShooter/src/helpers"
	"SpaceShooter/src/npcs"
	"SpaceShooter/src/player"
	"SpaceShooter/src/systems"
	"SpaceShooter/src/weapons"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/keyboard/keyboard"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	_ "image/png"
	"time"
)

//Convert this into something that loads the levels.

type Level struct {
	keys              []ebiten.Key
	enemies           []*npcs.Enemy
	soundEffectPlayer *audio.Player
	lastFire          time.Time
	playerBullets     []*weapons.Bullet
}

var (
	SCENENAME = "Level 1"
	PLAYER    = player.NewPLayer()
)

func (levelClass *Level) Init() {
	cX, cY := systems.WINDOWMANAGER.Center()

	systems.MUSICSYSTEM.LoadSong(systems.ASSETSYSTEM.Assets[SCENENAME].BackgroundMusic).PlaySong()
	PLAYER.Ship.SelectShip(1, 2)

	levelClass.enemies = append(levelClass.enemies, npcs.NewEnemy())

	PLAYER.XPos = cX
	PLAYER.YPos = cY

	levelClass.soundEffectPlayer, _ = audio.CurrentContext().NewPlayer(PLAYER.Ship.FireSound)

}

func (levelClass *Level) GetName() string {
	return "Level 1"
}

func NewLevelOne() *Level {
	g := &Level{}
	return g
}

func (levelClass *Level) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x00, 0x40, 0x80, 0xff})

	backgroundOP := &ebiten.DrawImageOptions{}
	backgroundOP.GeoM.Scale(2, 2)
	screen.DrawImage(systems.ASSETSYSTEM.Assets[SCENENAME].Images["Background"], backgroundOP)

	for i := 0; i < len(levelClass.playerBullets); i++ {
		op := &ebiten.DrawImageOptions{}
		op.Filter = ebiten.FilterLinear
		op.GeoM.Translate(levelClass.playerBullets[i].Xpos, levelClass.playerBullets[i].Ypos)
		screen.DrawImage(levelClass.playerBullets[i].Sprite, op)
	}

	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterLinear
	op.GeoM.Translate(PLAYER.XPos, PLAYER.YPos)
	screen.DrawImage(PLAYER.Ship.CurrentShipImage, op)

	for e := 0; e < len(levelClass.enemies); e++ {
		op := &ebiten.DrawImageOptions{}
		op.Filter = ebiten.FilterLinear
		op.GeoM.Translate(levelClass.enemies[e].PosX, levelClass.enemies[e].PosY)
		screen.DrawImage(levelClass.enemies[e].Image, op)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}
func RemoveIndex(s []*weapons.Bullet, index int) []*weapons.Bullet {
	return append(s[:index], s[index+1:]...)
}

func (levelClass *Level) Update() error {
	for i := 0; i < len(levelClass.playerBullets); i++ {
		levelClass.playerBullets[i].Ypos -= 10

		for e := 0; e < len(levelClass.enemies); e++ {

			if levelClass.enemies != nil && helpers.DistanceBetween(levelClass.enemies[e].PosX, levelClass.enemies[e].PosY, levelClass.playerBullets[i].Xpos, levelClass.playerBullets[i].Ypos) <= 40 {
				levelClass.enemies = append(levelClass.enemies[:e], levelClass.enemies[e+1:]...)
				levelClass.playerBullets = RemoveIndex(levelClass.playerBullets, i)
				break
			}

		}

		if len(levelClass.playerBullets) > 0 && levelClass.playerBullets[i].Ypos < 0 {
			levelClass.playerBullets = RemoveIndex(levelClass.playerBullets, i)
		}
	}

	levelClass.keys = inpututil.AppendPressedKeys(levelClass.keys[:0])
	for _, p := range levelClass.keys {
		_, ok := keyboard.KeyRect(p)

		if p.String() == "A" && (PLAYER.XPos > 0) {
			PLAYER.XPos -= 10
		}

		if p.String() == "D" && (PLAYER.XPos+(PLAYER.Ship.CurrentShipWidth) < float64(systems.WINDOWMANAGER.SCREENWIDTH)) {
			PLAYER.XPos += 10
		}

		if p.String() == "W" && (PLAYER.YPos > 0) {
			PLAYER.YPos -= 10
		}

		if p.String() == "S" && ((PLAYER.YPos + PLAYER.Ship.CurrentShipHeight) < float64(systems.WINDOWMANAGER.SCREENHEIGHT)) {
			PLAYER.YPos += 10
		}

		if p.String() == "Space" && !levelClass.soundEffectPlayer.IsPlaying() && (time.Now().Sub(levelClass.lastFire).Milliseconds() > PLAYER.Ship.FireRate) {

			bullet := weapons.NewBullet(systems.ASSETSYSTEM.Assets[SCENENAME].Images["LaserBullet"])
			bullet = bullet.SetCoordinates(PLAYER.XPos+(PLAYER.Ship.CurrentShipWidth/2)-(bullet.Width/2), PLAYER.YPos)
			levelClass.playerBullets = append(levelClass.playerBullets, bullet)
			levelClass.lastFire = time.Now()
			systems.MUSICSYSTEM.SetVolume(.50)
			levelClass.soundEffectPlayer.SetVolume(1)
			levelClass.soundEffectPlayer.Rewind()
			levelClass.soundEffectPlayer.Play()

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
