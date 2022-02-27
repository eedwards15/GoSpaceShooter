package scenes

import (
	"SpaceShooter/src/helpers"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
)

type MainMenu struct {
	titleArcadeFont font.Face
	Testing         string
}

func (mainMenuClass *MainMenu) GetName() string {
	return "Main Menu"
}

const (
	tileSize      = 32
	titleFontSize = fontSize * 1.5
	fontSize      = 24
	titleText     = "Hello World"
)

func NewMainMenu() *MainMenu {
	g := &MainMenu{}
	g.Init()
	return g
}

func (mainMenuClass *MainMenu) Init() {
	mainMenuClass.Testing = "Testing My Value"
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
	text.Draw(screen, titleText, mainMenuClass.titleArcadeFont, 1280-len(titleText)*fontSize, fontSize, color.White)
}

func (mainMenuClass *MainMenu) Update() error {
	return nil
}
