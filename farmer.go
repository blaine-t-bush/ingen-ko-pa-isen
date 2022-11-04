package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	FarmerSpeedMultiplier = 4
)

func CreateFarmer(img ebiten.Image) *Actor {
	w, h := img.Size()

	return &Actor{
		sprite: &Sprite{
			image:       &img,
			imageWidth:  float64(w),
			imageHeight: float64(h),
			pos:         &Coordinate{float64(screenWidth / 2), float64(screenHeight / 2)},
		},
	}
}
