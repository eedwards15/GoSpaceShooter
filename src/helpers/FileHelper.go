package helpers

import (
	"fmt"
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
