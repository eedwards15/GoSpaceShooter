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
	"github.com/hajimehoshi/ebiten/v2/examples/keyboard/keyboard"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"strconv"
	"time"
)

//Convert this into something that loads the levels.

type Level struct {
	keys              []ebiten.Key
	enemies           []*npcs.Enemy
	soundEffectPlayer *audio.Player
	lastFire          time.Time
	playerBullets     []*weapons.Bullet
	SCENENAME         string
	fxPlayer          *audio.Player
}

var (
	SCORE             = 0
	GOBAL_ASSETS      = "Global"
	LAST_SPAWN_TIME   = time.Now()
	PLAYER            = &player.Player{}
	TITLE_ARCADE_FONT font.Face
)

func (levelClass *Level) Init() {
	fmt.Println("INIT")
	levelClass.enemies = []*npcs.Enemy{}
	levelClass.SCENENAME = "Level 1"
	levelClass.playerBullets = []*weapons.Bullet{}
	SCORE = 0
	systems.MUSICSYSTEM.SetVolume(.50)
	cX, cY := systems.WINDOWMANAGER.Center()
	PLAYER = player.NewPLayer(cX, cY)
	PLAYER.IsDead = false

	soundEffect := systems.ASSETSYSTEM.Assets["Global"].SoundEffects["EnemyExplosion"]
	levelClass.fxPlayer, _ = audio.CurrentContext().NewPlayer(soundEffect)

	systems.MUSICSYSTEM.LoadSong(systems.ASSETSYSTEM.Assets[levelClass.SCENENAME].BackgroundMusic).PlaySong()
	PLAYER.Ship.SelectShip(1, 2)
	levelClass.enemies = append(levelClass.enemies, npcs.NewEnemy(float64(systems.WINDOWMANAGER.SCREENWIDTH/2), 0))
	levelClass.soundEffectPlayer, _ = audio.CurrentContext().NewPlayer(PLAYER.Ship.FireSound)

	tt, err := opentype.Parse(*helpers.LoadFile("assets/fonts/arcades/Arcades.ttf"))
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	TITLE_ARCADE_FONT, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    18,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

}

func (levelClass *Level) GetName() string {
	return "Level 1"
}

func NewLevel() *Level {
	g := &Level{}
	return g
}

func (levelClass *Level) Draw(screen *ebiten.Image) {
	backgroundOP := &ebiten.DrawImageOptions{}
	backgroundOP.GeoM.Scale(2, 2)
	screen.DrawImage(systems.ASSETSYSTEM.Assets[levelClass.SCENENAME].Images["Background"], backgroundOP)

	for i := 0; i < len(levelClass.playerBullets); i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(levelClass.playerBullets[i].Xpos, levelClass.playerBullets[i].Ypos)
		screen.DrawImage(levelClass.playerBullets[i].Sprite, op)
	}

	for e := 0; e < len(levelClass.enemies); e++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(levelClass.enemies[e].PosX, levelClass.enemies[e].PosY)
		screen.DrawImage(levelClass.enemies[e].Image, op)
	}

	PLAYER.Draw(screen)
	text.Draw(screen, "Score: "+strconv.Itoa(SCORE), TITLE_ARCADE_FONT, 20, 20, color.White)

	//ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}
func RemoveIndex(s []*weapons.Bullet, index int) []*weapons.Bullet {
	return append(s[:index], s[index+1:]...)
}

func (levelClass *Level) Update() error {
	if PLAYER.IsDead {
		return nil
	}

	for e := 0; e < len(levelClass.enemies); e++ {

		if helpers.DistanceBetween(PLAYER.XPos, PLAYER.YPos, levelClass.enemies[e].PosX, levelClass.enemies[e].PosY) <= 50 {
			PLAYER.IsDead = true
			systems.SCENEMANAGER.Push(NewGameOver())
			return nil
		}

	}

	for i := 0; i < len(levelClass.playerBullets); i++ {
		levelClass.playerBullets[i].Ypos -= 10

		removeBullet := false
		//Enemy Loop

		//EnemyExplosion
		for e := 0; e < len(levelClass.enemies); e++ {
			//Check to see if any bullets hit any of the enemies.
			if levelClass.enemies != nil && helpers.DistanceBetween(levelClass.enemies[e].PosX+float64(levelClass.enemies[e].Width/2), levelClass.enemies[e].PosY, levelClass.playerBullets[i].Xpos, levelClass.playerBullets[i].Ypos) <= 50 {
				levelClass.enemies[e].Dead = true
				removeBullet = true
				SCORE += 10
				levelClass.fxPlayer.Rewind()
				levelClass.fxPlayer.Play()
				break
			}

		}

		//Clean up dead enemies
		newEnemyList := []*npcs.Enemy{}
		for e := 0; e < len(levelClass.enemies); e++ {
			if levelClass.enemies[e].Dead {
				continue
			}

			newEnemyList = append(newEnemyList, levelClass.enemies[e])
		}

		levelClass.enemies = newEnemyList

		//If Bullet Gets out of screen range remove.
		if removeBullet == true || len(levelClass.playerBullets) > 0 && levelClass.playerBullets[i].Ypos < 0 {
			levelClass.playerBullets = RemoveIndex(levelClass.playerBullets, i)
		}
	}

	//Create New Enemy
	if time.Now().Sub(LAST_SPAWN_TIME).Seconds() > 2 {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		x := r1.Intn(systems.WINDOWMANAGER.SCREENWIDTH - 50)
		levelClass.enemies = append(levelClass.enemies, npcs.NewEnemy(float64(x), 0))
		LAST_SPAWN_TIME = time.Now()
	}

	//Enemy Movement
	for e := 0; e < len(levelClass.enemies); e++ {
		levelClass.enemies[e].PosY += 5

		if levelClass.enemies[e].PosY > float64(systems.WINDOWMANAGER.SCREENHEIGHT) {
			levelClass.enemies[e].PosY = 0
		}
	}

	levelClass.keys = inpututil.AppendPressedKeys(levelClass.keys[:0])
	for _, p := range levelClass.keys {
		_, ok := keyboard.KeyRect(p)

		if p.String() == "A" && (PLAYER.XPos > 0) {
			PLAYER.MoveX(-10)
		}

		if p.String() == "D" && (PLAYER.XPos+(PLAYER.Ship.CurrentShipWidth) < float64(systems.WINDOWMANAGER.SCREENWIDTH)) {
			PLAYER.MoveX(10)
		}

		if p.String() == "W" && (PLAYER.YPos > 0) {
			PLAYER.MoveY(-10)
		}

		if p.String() == "S" && ((PLAYER.YPos + PLAYER.Ship.CurrentShipHeight) < float64(systems.WINDOWMANAGER.SCREENHEIGHT)) {
			PLAYER.MoveY(10)
		}

		if p.String() == "Space" && (time.Now().Sub(levelClass.lastFire).Milliseconds() > PLAYER.Ship.FireRate) {
			//Look into Command Pattern.
			bullet := weapons.NewBullet(systems.ASSETSYSTEM.Assets[GOBAL_ASSETS].Images["LaserBullet"])
			bullet = bullet.SetCoordinates(PLAYER.XPos+(PLAYER.Ship.CurrentShipWidth/2)-(bullet.Width/2), PLAYER.YPos)
			levelClass.playerBullets = append(levelClass.playerBullets, bullet)
			levelClass.lastFire = time.Now()
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
