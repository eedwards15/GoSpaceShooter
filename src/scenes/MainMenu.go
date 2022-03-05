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
	"log"
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
	fontSize      = 24
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
	tt, err := opentype.Parse(*helpers.LoadFile("./assets/fonts/arcades/Arcades.ttf"))
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mainMenuClass.titleArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    titleFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

}

func (mainMenuClass *MainMenu) Draw(screen *ebiten.Image) {
	if mainMenuClass.titleArcadeFont == nil {
		return
	}
	text.Draw(screen, titleText, mainMenuClass.titleArcadeFont, (systems.WINDOWMANAGER.SCREENWIDTH/2)-(len(titleText)/2)*fontSize, (systems.WINDOWMANAGER.SCREENHEIGHT / 2), color.White)
	text.Draw(screen, instructions, mainMenuClass.titleArcadeFont, (systems.WINDOWMANAGER.SCREENWIDTH/2)-(len(instructions)/2)*fontSize, (systems.WINDOWMANAGER.SCREENHEIGHT/2)+fontSize*3, color.White)

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
