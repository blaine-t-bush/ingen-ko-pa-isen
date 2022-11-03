package main

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewImageFromFilePath(filePath string) *ebiten.Image {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	image, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(image)
}
