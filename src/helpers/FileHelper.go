package helpers

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"os"
)

func LoadFile(location string) *[]byte {
	_, error := os.Stat(location)
	if error != nil {
		fmt.Println("Location Not Found")
		panic("Could not find file")
	}

	fileData, fileError := os.ReadFile(location)
	if fileError != nil {
		fmt.Println("Can not read file")
		panic("Can not open file")
	}

	return &fileData
}

func OpenImage(location string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(location)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
