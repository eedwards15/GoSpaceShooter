package helpers

import (
	_ "github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Helpers struct {
}

func CenterTextXPos(value string, fontUsed font.Face, screenWidth int) int {
	return (screenWidth - (font.MeasureString(fontUsed, value).Floor())) / 2
}
