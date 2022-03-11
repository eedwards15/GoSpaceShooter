package scenes

import (
	"SpaceShooter/assets"
	"SpaceShooter/src/helpers"
	"SpaceShooter/src/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/keyboard/keyboard"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
	"strconv"
)

type GameOver struct {
	titleArcadeFont font.Face
	keys            []ebiten.Key
	SCENENAME       string
}

func (gameOverClass *GameOver) GetName() string {
	return "Game Over"
}

func NewGameOver() *GameOver {
	g := &GameOver{}

	return g
}

func (gameOverClass *GameOver) Init() {
	gameOverClass.SCENENAME = "GameOver"
	systems.MUSICSYSTEM.LoadSong(systems.ASSETSYSTEM.Assets[gameOverClass.SCENENAME].BackgroundMusic).PlaySong()
	f, _ := assets.AssetsFileSystem.ReadFile("fonts/arcades/Arcades.ttf")
	tt, err := opentype.Parse(f)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	gameOverClass.titleArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    50,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

}

func (gameOverClass *GameOver) Draw(screen *ebiten.Image) {
	if gameOverClass.titleArcadeFont == nil {
		return
	}
	text.Draw(screen, "Your Score: "+strconv.Itoa(SCORE), gameOverClass.titleArcadeFont, helpers.CenterTextXPos("Your Score: "+strconv.Itoa(SCORE), gameOverClass.titleArcadeFont, systems.WINDOWMANAGER.SCREENWIDTH), 150, color.White)
	text.Draw(screen, "Game Over", gameOverClass.titleArcadeFont, helpers.CenterTextXPos("Game Over", gameOverClass.titleArcadeFont, systems.WINDOWMANAGER.SCREENWIDTH), 350, color.White)
	text.Draw(screen, instructions, gameOverClass.titleArcadeFont, helpers.CenterTextXPos(instructions, gameOverClass.titleArcadeFont, systems.WINDOWMANAGER.SCREENWIDTH), 550, color.White)
}

func (gameOverClass *GameOver) Update() error {
	gameOverClass.keys = inpututil.AppendPressedKeys(gameOverClass.keys[:0])

	for _, p := range gameOverClass.keys {
		_, ok := keyboard.KeyRect(p)

		if p.String() == "Enter" {
			systems.SCENEMANAGER.Pop()
		}

		if !ok {
			continue
		}

	}

	return nil
}
