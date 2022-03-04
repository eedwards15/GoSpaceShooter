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
	ckX, bky int
)

func (levelOneClass *Level) Init() {
	systems.MUSICSYSTEM.LoadSong(systems.ASSETSYSTEM.LevelOne.BackgroundMusic).PlaySong()
	player.PLAYER.Ship.SelectShip(1, 2)
	levelOneClass.enemies = append(levelOneClass.enemies, npcs.NewEnemy())

	cX, cY := systems.WINDOWMANAGER.Center()
	player.PLAYER.XPos = cX
	player.PLAYER.YPos = cY

	levelOneClass.soundEffectPlayer, _ = audio.CurrentContext().NewPlayer(player.PLAYER.Ship.FireSound)

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

	backgroundOP := &ebiten.DrawImageOptions{}
	backgroundOP.GeoM.Scale(2, 2)
	screen.DrawImage(systems.ASSETSYSTEM.LevelOne.Images["Background"], backgroundOP)

	for i := 0; i < len(levelOneClass.playerBullets); i++ {
		op := &ebiten.DrawImageOptions{}
		op.Filter = ebiten.FilterLinear
		op.GeoM.Translate(levelOneClass.playerBullets[i].Xpos, levelOneClass.playerBullets[i].Ypos)
		screen.DrawImage(levelOneClass.playerBullets[i].Sprite, op)
	}

	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterLinear
	op.GeoM.Translate(player.PLAYER.XPos, player.PLAYER.YPos)
	screen.DrawImage(player.PLAYER.Ship.CurrentShipImage, op)

	for e := 0; e < len(levelOneClass.enemies); e++ {
		op := &ebiten.DrawImageOptions{}
		op.Filter = ebiten.FilterLinear
		op.GeoM.Translate(levelOneClass.enemies[e].PosX, levelOneClass.enemies[e].PosY)
		screen.DrawImage(levelOneClass.enemies[e].Image, op)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}
func RemoveIndex(s []*weapons.Bullet, index int) []*weapons.Bullet {
	return append(s[:index], s[index+1:]...)
}

func (levelOneClass *Level) Update() error {
	for i := 0; i < len(levelOneClass.playerBullets); i++ {
		levelOneClass.playerBullets[i].Ypos -= 10

		for e := 0; e < len(levelOneClass.enemies); e++ {

			if levelOneClass.enemies != nil && helpers.DistanceBetween(levelOneClass.enemies[e].PosX, levelOneClass.enemies[e].PosY, levelOneClass.playerBullets[i].Xpos, levelOneClass.playerBullets[i].Ypos) <= 40 {
				levelOneClass.enemies = append(levelOneClass.enemies[:e], levelOneClass.enemies[e+1:]...)
				levelOneClass.playerBullets = RemoveIndex(levelOneClass.playerBullets, i)
				break
			}

		}

		if len(levelOneClass.playerBullets) > 0 && levelOneClass.playerBullets[i].Ypos < 0 {
			levelOneClass.playerBullets = RemoveIndex(levelOneClass.playerBullets, i)
		}
	}

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

		if p.String() == "Space" && !levelOneClass.soundEffectPlayer.IsPlaying() && (time.Now().Sub(levelOneClass.lastFire).Milliseconds() > player.PLAYER.Ship.FireRate) {

			bullet := weapons.NewBullet(systems.ASSETSYSTEM.LevelOne.Images["LaserBullet"])
			bullet = bullet.SetCoordinates(player.PLAYER.XPos+(player.PLAYER.Ship.CurrentShipWidth/2)-(bullet.Width/2), player.PLAYER.YPos)
			levelOneClass.playerBullets = append(levelOneClass.playerBullets, bullet)
			levelOneClass.lastFire = time.Now()
			systems.MUSICSYSTEM.SetVolume(.50)
			levelOneClass.soundEffectPlayer.SetVolume(1)
			levelOneClass.soundEffectPlayer.Rewind()
			levelOneClass.soundEffectPlayer.Play()

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
