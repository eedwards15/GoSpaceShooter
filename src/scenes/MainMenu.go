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
	Testing         string
	keys            []ebiten.Key
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
	return g
}

func (mainMenuClass *MainMenu) Init() {

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
	mainMenuClass.keys = inpututil.AppendPressedKeys(mainMenuClass.keys[:0])

	for _, p := range mainMenuClass.keys {
		_, ok := keyboard.KeyRect(p)

		if p.String() == "Enter" {
			systems.SCENEMANAGER.Push(NewLevelOne())
		}

		if !ok {
			continue
		}

	}

	return nil
}
