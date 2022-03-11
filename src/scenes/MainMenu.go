package scenes

import (
	"SpaceShooter/src/helpers"
	"SpaceShooter/src/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/keyboard/keyboard"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
)

type MainMenu struct {
	titleArcadeFont font.Face
	keys            []ebiten.Key
	SCENENAME       string
}

func (mainMenuClass *MainMenu) GetName() string {
	return "Main Menu"
}

const (
	tileSize      = 32
	titleFontSize = fontSize * 1.5
	fontSize      = 100
	titleText     = "Space Shooter"
	instructions  = "Press Enter To Play"
)

func NewMainMenu() *MainMenu {
	g := &MainMenu{}

	return g
}

func (mainMenuClass *MainMenu) Init() {
	mainMenuClass.SCENENAME = "MainMenu"
	systems.MUSICSYSTEM.LoadSong(systems.ASSETSYSTEM.Assets[mainMenuClass.SCENENAME].BackgroundMusic).PlaySong()
	mainMenuClass.titleArcadeFont, _ = opentype.NewFace(systems.ASSETSYSTEM.Assets["Global"].Fonts["Arcades"], &opentype.FaceOptions{
		Size:    100,
		DPI:     72,
		Hinting: font.HintingFull,
	})

}

func (mainMenuClass *MainMenu) Draw(screen *ebiten.Image) {
	if mainMenuClass.titleArcadeFont == nil {
		return
	}
	backgroundOP := &ebiten.DrawImageOptions{}
	backgroundOP.GeoM.Scale(2, 2)
	screen.DrawImage(systems.ASSETSYSTEM.Assets["Level 1"].Images["Background"], backgroundOP)

	text.Draw(screen, titleText, mainMenuClass.titleArcadeFont, helpers.CenterTextXPos(titleText, mainMenuClass.titleArcadeFont, systems.WINDOWMANAGER.SCREENWIDTH), 200, color.White)

	text.Draw(screen, "Play", mainMenuClass.titleArcadeFont, helpers.CenterTextXPos("Play", mainMenuClass.titleArcadeFont, systems.WINDOWMANAGER.SCREENWIDTH), 400, color.White)

	text.Draw(screen, "Exit", mainMenuClass.titleArcadeFont, helpers.CenterTextXPos("Exit", mainMenuClass.titleArcadeFont, systems.WINDOWMANAGER.SCREENWIDTH), 600, color.White)

}

func (mainMenuClass *MainMenu) Update() error {
	mainMenuClass.keys = inpututil.AppendPressedKeys(mainMenuClass.keys[:0])

	for _, p := range mainMenuClass.keys {
		_, ok := keyboard.KeyRect(p)

		if p.String() == "Enter" {
			systems.SCENEMANAGER.Push(NewLevel())
		}

		if !ok {
			continue
		}

	}

	return nil
}
