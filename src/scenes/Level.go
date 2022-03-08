package scenes

import (
	"SpaceShooter/assets"
	"SpaceShooter/src/helpers"
	"SpaceShooter/src/npcs"
	"SpaceShooter/src/player"
	"SpaceShooter/src/systems"
	"SpaceShooter/src/weapons"
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
	keys                   []ebiten.Key
	enemies                []npcs.IEnemy
	soundEffectPlayer      *audio.Player
	soundEffectPlayerDeath *audio.Player
	lastFire               time.Time
	playerBullets          []*weapons.Bullet
	EnemyBullets           []*weapons.Bullet
	SCENENAME              string
	fxPlayer               *audio.Player
}

var (
	SCORE             = 0
	GOBAL_ASSETS      = "Global"
	LAST_SPAWN_TIME   = time.Now()
	PLAYER            = &player.Player{}
	TITLE_ARCADE_FONT font.Face
)

func (levelClass *Level) Init() {
	levelClass.enemies = []npcs.IEnemy{}
	levelClass.SCENENAME = "Level 1"
	levelClass.playerBullets = []*weapons.Bullet{}
	levelClass.EnemyBullets = []*weapons.Bullet{}

	SCORE = 0
	systems.MUSICSYSTEM.SetVolume(.50)
	cX, cY := systems.WINDOWMANAGER.Center()
	PLAYER = player.NewPLayer(cX, cY)
	PLAYER.IsDead = false

	soundEffect := systems.ASSETSYSTEM.Assets["Global"].SoundEffects["EnemyExplosion"]
	levelClass.fxPlayer, _ = audio.CurrentContext().NewPlayer(soundEffect)
	levelClass.soundEffectPlayerDeath, _ = audio.CurrentContext().NewPlayer(systems.ASSETSYSTEM.Assets["Global"].SoundEffects["PlayerDeath"])

	systems.MUSICSYSTEM.LoadSong(systems.ASSETSYSTEM.Assets[levelClass.SCENENAME].BackgroundMusic).PlaySong()
	PLAYER.Ship.SelectShip(1, 2)

	levelClass.enemies = append(levelClass.enemies, npcs.NewWeakEnemy(float64(systems.WINDOWMANAGER.SCREENWIDTH/2), 0))
	levelClass.soundEffectPlayer, _ = audio.CurrentContext().NewPlayer(PLAYER.Ship.FireSound)
	f, _ := assets.AssetsFileSystem.ReadFile("fonts/arcades/Arcades.ttf")
	tt, err := opentype.Parse(f)
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

	for i := 0; i < len(levelClass.EnemyBullets); i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(levelClass.EnemyBullets[i].Xpos, levelClass.EnemyBullets[i].Ypos)
		screen.DrawImage(levelClass.EnemyBullets[i].Sprite, op)
	}

	for e := 0; e < len(levelClass.enemies); e++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(levelClass.enemies[e].GetPosX(), levelClass.enemies[e].GetPosY())
		screen.DrawImage(levelClass.enemies[e].GetImage(), op)
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

		//Check for collision with player.
		if helpers.DistanceBetween(PLAYER.XPos, PLAYER.YPos, levelClass.enemies[e].GetPosX(), levelClass.enemies[e].GetPosY()) <= 50 {
			PLAYER.IsDead = true
			systems.SCENEMANAGER.Push(NewGameOver())
			levelClass.soundEffectPlayerDeath.Rewind()
			levelClass.soundEffectPlayerDeath.Play()
			return nil
		}

		if levelClass.enemies[e].CanShoot() {
			//Fire Bullets
			ebullet := levelClass.enemies[e].Fire() // Will return nil if they can't fire yet.
			if ebullet != nil {
				ebullet.SetCoordinates(levelClass.enemies[e].GetPosX()+float64(levelClass.enemies[e].GetWidth()/2)-(ebullet.Width/2), levelClass.enemies[e].GetPosY())
				levelClass.EnemyBullets = append(levelClass.EnemyBullets, ebullet)
			}
		}
	}

	for i := 0; i < len(levelClass.playerBullets); i++ {
		levelClass.playerBullets[i].Ypos -= 10
		removeBullet := false

		//EnemyExplosion
		for e := 0; e < len(levelClass.enemies); e++ {
			//Check to see if any bullets hit any of the enemies.
			if levelClass.enemies != nil && helpers.DistanceBetween(levelClass.enemies[e].GetPosX()+float64(levelClass.enemies[e].GetWidth()/2), levelClass.enemies[e].GetPosY(), levelClass.playerBullets[i].Xpos, levelClass.playerBullets[i].Ypos) <= 50 {
				levelClass.enemies[e].TakeDamage()
				removeBullet = true

				if levelClass.enemies[e].IsDead() {
					SCORE += levelClass.enemies[e].GetScoreAmount()
					levelClass.fxPlayer.Rewind()
					levelClass.fxPlayer.Play()
				}

				break
			}

		}

		//Clean up dead enemies
		newEnemyList := []npcs.IEnemy{}
		for e := 0; e < len(levelClass.enemies); e++ {
			if levelClass.enemies[e].IsDead() {
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

	for i := 0; i < len(levelClass.EnemyBullets); i++ {
		levelClass.EnemyBullets[i].Ypos += 10

		if helpers.DistanceBetween(PLAYER.XPos+float64(PLAYER.Ship.CurrentShipWidth/2), PLAYER.YPos, levelClass.EnemyBullets[i].Xpos, levelClass.EnemyBullets[i].Ypos) <= 30 {
			PLAYER.IsDead = true
			systems.SCENEMANAGER.Push(NewGameOver())
			levelClass.soundEffectPlayerDeath.Rewind()
			levelClass.soundEffectPlayerDeath.Play()
			return nil
		}

		if len(levelClass.EnemyBullets) > 0 && levelClass.EnemyBullets[i].Ypos > float64(systems.WINDOWMANAGER.SCREENHEIGHT) {
			levelClass.EnemyBullets = RemoveIndex(levelClass.EnemyBullets, i)
		}
	}

	//Create New WeakEnemy
	//Replace With factory
	if time.Now().Sub(LAST_SPAWN_TIME).Seconds() > 2 {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		x := r1.Intn(systems.WINDOWMANAGER.SCREENWIDTH - 100)

		newEnemey := npcs.SpawnNewEnemy(float64(x), 0)
		levelClass.enemies = append(levelClass.enemies, newEnemey)

		LAST_SPAWN_TIME = time.Now()
	}

	//WeakEnemy Movement
	for e := 0; e < len(levelClass.enemies); e++ {
		currentLocationY := levelClass.enemies[e].GetPosY()
		levelClass.enemies[e].SetPosY(currentLocationY + 5)

		//Moves the WeakEnemy Back To the Top of the screen
		if levelClass.enemies[e].GetPosY() > float64(systems.WINDOWMANAGER.SCREENHEIGHT) {
			levelClass.enemies[e].SetPosY(0)
		}
	}

	//INPUTs
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
